import tornado.options
import tornado.template
import json
import locale
import os.path
import datetime
from collections import defaultdict

def pct_of(a, b):
    delta = float(a) / (b) * 100
    return "%0.2f%%" % delta

def delta_ms(a, b):
    delta = a - b
    if delta > 0:
        txt = "+" + locale.format("%d", delta, grouping=True)
    else:
        txt = locale.format("%d", delta, grouping=True)
    return '<span class="%s">%s<small>ms</small></span>' % (delta_ms_css(delta), txt)

def delta(a, b, up=True):
    assert isinstance(up, bool)
    delta = a - b
    pct = 0
    if b != 0:
        pct = ((float(a) / b) * 100) - 100
    if delta > 0:
        txt = "+" + locale.format("%d", delta, grouping=True)
    else:
        txt = locale.format("%d", delta, grouping=True)
    return '<span class="%s">%s</span>' % (delta_pct_css(pct, up), txt)

def delta_pct_css(delta, up):
    css = ""
    if delta > 10:
        css = "down down-10"
    elif delta > 5:
        css ="down down-5"
    elif delta > 0:
        css = "down"
    elif delta < -10:
        css = "up up-10"
    elif delta < -5:
        css = "up up-5"
    else:
        css = "up"
    if not up:
        if 'up' in css:
            return css.replace('up', 'down')
        return css.replace('down', 'up')
        
    return css

def delta_ms_css(delta):
    css = ""
    if delta > 250:
        css = "up-10"
    elif delta > 100:
        css ="up-5"
    elif delta > 0:
        css = "up"
    elif delta < -250:
        css = "down-10"
    elif delta < -100:
        css = "down-5"
    else:
        css = "down"
    return css

def load_file(filename):
    data = defaultdict(dict)
    if not os.path.exists(filename):
        return data
    with open(filename, 'r') as f:
        for line in f:
            row = json.loads(line)
            data[row['code']] = row
    return data

def file_date(filename):
    if not filename:
        return None
    date_str, _ = os.path.splitext(os.path.basename(filename))
    date_str = date_str.split("_")[-1] # YYYY-MM-dd
    return datetime.datetime.strptime(date_str, "%Y-%m-%d")

def run(current_file, previous_file):
    report_dt = file_date(current_file)
    previous_dt = file_date(previous_file)
    loader = tornado.template.Loader(".", autoescape=None)

    current_data = load_file(current_file)
    assert current_data
    previous_data = load_file(previous_file)

    t = loader.load("sign_template.html")
    output = t.generate(
        locale=locale,
        report_dt=report_dt,
        previous_dt=previous_dt,
        current_data = current_data,
        previous_data = previous_data,
        pct_of=pct_of,
        delta_ms=delta_ms,
        delta=delta,
        categories=(
            ("ParkingSign", "Parking Signs"),
            ("StreetCleaning", "Street Cleaning Signs"),
            ("InformationSign", "Informational Signs"),
            ("BusInformation", "Bus Related Signs"),
            ("SpecialInterestParking", "Special Interest Signs"),
            ("OtherRegulationSign", "Other Regulations"),
            ("UnknownSign", "Unknown Signs"),
        ),
    )
    
    with open("data/signs_%s.html" % report_dt.strftime("%Y-%m-%d"), 'w') as f:
        f.write(output)
    

if __name__ == "__main__":
    tornado.options.define("current_file", type=str, default="")
    tornado.options.define("previous_file", type=str, default="")
    tornado.options.parse_command_line()

    locale.setlocale(locale.LC_ALL, 'en_US')
    o = tornado.options.options
    assert os.path.exists(o.current_file)
    if o.previous_file:
        assert os.path.exists(o.previous_file)
    run(o.current_file, o.previous_file)