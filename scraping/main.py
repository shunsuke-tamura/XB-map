import requests
import googlemaps
from bs4 import BeautifulSoup
from time import sleep

def get_googlemaps_geocoding(query):
  '''
  Args:
      query(str): 住所
  Returns:
      lat, lng(float): 緯度、経度
  '''
  googleapikey = "AIzaSyCqWh2yeQ_Pu2QQa5FV_eAo_rhESOueEGE"
  gmaps = googlemaps.Client(key=googleapikey)
  gmap_list = gmaps.geocode(query)
  try:
    ll = gmap_list[0]['geometry']['location']
    lat, lng = ll['lat'], ll['lng']
  except:
    lat = 0
    lng = 0

  return lat, lng

for i in range(77,78,1):
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

      # adress(住所)の座標が見つからなければ．shop(店舗名)の座標を入れる
      adress = shop_info.find_all("p", attrs={"class": "add"})[0].text
      lat, lng = get_googlemaps_geocoding(adress)
      if lat == 0 and lng == 0:
        lat, lng = get_googlemaps_geocoding(shop)

      # 2人対戦対応店，冒険の書マシン販売対応店，冒険の書セット取扱店
      if not shop_info.find_all("li", attrs={"class", "type01"}):
        two_player = ""
      else:
        two_player = shop_info.find_all("li", attrs={"class", "type01"})[0].text

      if not shop_info.find_all("li", attrs={"class", "type02"}):
        adventure_book_machine_sale = ""
      else:
        adventure_book_machine_sale = shop_info.find_all("li", attrs={"class", "type02"})[0].text
        
      if not shop_info.find_all("li", attrs={"class", "type03"}):
        adventure_book_set = ""
      else:
        adventure_book_set = shop_info.find_all("li", attrs={"class", "type03"})[0].text


      print(shop)
      print(adress)
      print(lat, lng)
      print(two_player)
      print(adventure_book_machine_sale)
      print(adventure_book_set)

  sleep(3)