from bs4 import BeautifulSoup
from requests import post, get, Response
from json import loads
from flask import Flask
app = Flask(__name__)

def GetMacSerialInfo(serial: str):
    postBody = b'["' + serial.encode('utf-8') + b'"]'
    resp = post("http://macserial", data=postBody)
    info = loads(resp.content)
    if info[serial] == "ERR":
        raise ValueError("Invalid serial or request for serial.")

    return info[serial]

def GetEveryMacName(year, model):
    url = "https://everymac.com/ultimate-mac-lookup/"
    params = {"search_keywords": model}

    headers = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36'}
    resp = get(url, params, headers=headers)
    
    soup = BeautifulSoup(resp.text, features="html.parser")
    text = soup.get_text()
    lines = text.splitlines()
    lines = [i.strip() for i in lines if i.strip() != ""]
    nextIter = False
    name = ""
    for i in lines:
        if nextIter:
            name = i
            break
        if "Showing" in i:
            nextIter = True
    if name == "":
        raise ValueError("Could not find corresponding modelname")
    return name

def SerialToModelName(serial):
    info = GetMacSerialInfo(serial)
    modelName = GetEveryMacName(info['Year']['Value'], info['Model']['Value'])
    return modelName

@app.route("/<string:serial>")
def SerialLookup(serial):
    try:
        m = SerialToModelName(serial)
        return m
    except Exception as e:
        return  e.args[0]

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=80)

    