#!/usr/local/bin/python
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

def delta(a, b, flip=False):
    assert isinstance(flip, bool)
    delta = a - b
    pct = 0
    if b != 0:
        pct = ((float(a) / b) * 100) - 100
    elif a > 0:
        pct = 100
    if delta > 0:
        txt = "+" + locale.format("%d", delta, grouping=True)
    else:
        txt = locale.format("%d", delta, grouping=True)
    return '<span class="%s">%s</span>' % (delta_pct_css(pct, flip), txt)

def delta_pct_css(delta_pct, flip):
    css = ""
    if delta_pct == 0:
        return "nochange"
    if delta_pct > 10:
        css = "up up-10"
    elif delta_pct > 5:
        css ="up up-5"
    elif delta_pct > 0:
        css = "up"
    elif delta_pct < -10:
        css = "down down-10"
    elif delta_pct < -5:
        css = "down down-5"
    else:
        css = "down"
    if flip:
        if 'up' in css:
            return css.replace('up', 'down')
        return css.replace('down', 'up')
        
    return css

def load_file(filename):
    data = defaultdict(dict)
    if not os.path.exists(filename):
        return data
    with open(filename, 'r') as f:
        for line in f:
            row = json.loads(line)
            assert 'type' in row, "%r" % row
            data[row['code']] = row
    return data

def file_date(filename):
    if not filename:
        return None
    date_str, _ = os.path.splitext(os.path.basename(filename))
    date_str = date_str.split("_")[-1] # YYYY-MM-dd
    return datetime.datetime.strptime(date_str, "%Y-%m-%d")

def run(current_file, previous_file):
    current_dt = file_date(current_file)
    previous_dt = file_date(previous_file)
    loader = tornado.template.Loader(".", autoescape=None)

    current_data = load_file(current_file)
    assert current_data
    previous_data = load_file(previous_file)

    t = loader.load("sign_template.html")
    output = t.generate(
        locale=locale,
        current_dt=current_dt,
        previous_dt=previous_dt,
        current_data = current_data,
        previous_data = previous_data,
        pct_of=pct_of,
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
    
    with open("data/nyc_dot_signs_%s.html" % current_dt.strftime("%Y-%m-%d"), 'w') as f:
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