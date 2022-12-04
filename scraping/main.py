import os
import googlemaps
import mysql.connector
import requests
from time import sleep
from bs4 import BeautifulSoup
from dotenv import load_dotenv


def get_googlemaps_geocoding(query):
  '''
  Args:
      query(str): 住所
  Returns:
      lat, lng(float): 緯度、経度
  '''
  # .envファイルの内容を読み込む
  load_dotenv()
  googleapikey = os.environ["GOOGLE_API_KEY"]
  gmaps = googlemaps.Client(key=googleapikey)
  gmap_list = gmaps.geocode(query)
  try:
    ll = gmap_list[0]['geometry']['location']
    lat, lng = ll['lat'], ll['lng']
  except:
    lat = 0
    lng = 0

  return lat, lng

def main():
  for i in range(1,100,1):
    # 全ページのhtmlデータを取得
    res = requests.get("https://www.dqdai-xb.jp/shop/search?&page=" + str(i))
    
    # .contentとすることでhtmlとして渡せる
    shops_list = BeautifulSoup(res.content, "html.parser")

    # urlが存在しないところまでfor文を回したらbreak
    if not shops_list.body.div.section.div.find_all("ul", attrs={"class": "shopList"}):
      print("end")
      break

    # お店の名前，住所，座標，お店のタイプが得られる
    for shop_list in shops_list.body.div.section.div.find_all("ul", attrs={"class": "shopList"}):
      for shop_info in shop_list.find_all("li", attrs={"class": "shop"}):
        shop = shop_info.find_all("dt", attrs={"class": "name"})[0].text
        adress = shop_info.find_all("p", attrs={"class": "add"})[0].text

        # dbをNoneと先に定義しておく
        db = None
        # dbと接続
        try:
          db = mysql.connector.connect(
            user="user", # ユーザー名
            password="password", # パスワード
            host="db_xb", # ホスト名(IPアドレス
            db="xb-map" # データベース名
          )
          if db.is_connected:
            cursor = db.cursor()
            # データベースに同じ名前の店舗があるか確認
            cursor.execute("SELECT EXISTS(SELECT name, adress FROM shops WHERE name = %s AND adress = %s)", (shop, adress))
            # 無ければデータを挿入する処理を続ける
            if cursor.fetchone()[0] == 0:
              # adress(住所)の座標が見つからなければ．shop(店舗名)の座標を入れる
              lat, lng = get_googlemaps_geocoding(adress)
              if lat == 0 and lng == 0:
                lat, lng = get_googlemaps_geocoding(shop)

              # 2人対戦対応店，冒険の書マシン販売対応店，冒険の書セット取扱店
              if not shop_info.find_all("li", attrs={"class", "type01"}):
                two_player = None
              else:
                two_player = shop_info.find_all("li", attrs={"class", "type01"})[0].text

              if not shop_info.find_all("li", attrs={"class", "type02"}):
                adventure_book_machine_sale = None
              else:
                adventure_book_machine_sale = shop_info.find_all("li", attrs={"class", "type02"})[0].text
                
              if not shop_info.find_all("li", attrs={"class", "type03"}):
                adventure_book_set = None
              else:
                adventure_book_set = shop_info.find_all("li", attrs={"class", "type03"})[0].text

              try:
                if db.is_connected:
                  # データをデータベースに追加
                  cursor.execute("""INSERT INTO shops(name, adress, lat, lng, type01, type02, type03)
                                    VALUE(%s, %s, %s, %s, %s, %s, %s)
                                """, 
                                (shop,
                                 adress,
                                 lat, 
                                 lng, 
                                 two_player, 
                                 adventure_book_machine_sale, 
                                 adventure_book_set))
                  # データベースに反映
                  db.commit()
              except Exception as e:
                print(e)
            else:
              print(shop)
              print(adress)
        except Exception as e:
          print(e)
        finally:
          if db is not None and db.is_connected():
            cursor.close()
            db.close()
    
    # 連続してサイトにアクセスするといけないので5秒空ける
    print(i)
    sleep(5)

if __name__ == '__main__':
  main()