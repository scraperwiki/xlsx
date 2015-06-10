#! /usr/bin/env python

# Requires freetype-py
# https://github.com/rougier/freetype-py
# https://pypi.python.org/pypi/freetype-py/0.4.2

from freetype import Face
import string

face = Face("/usr/share/fonts/truetype/msttcorefonts/arial.ttf")

font_size_pt = 11

face.set_char_size(font_size_pt << 6)

def width(c):
	face.load_char(c)
	return face.glyph.advance.x >> 6

def iterate_glyphs():
	code, result = face.get_first_char()
	while result != 0:
		face.load_char(unichr(code))
		yield code, face.glyph.advance.x
		code, result = face.get_next_char(code, 0)
		# print unichr(code), result

for code, advance in iterate_glyphs():
	print u"{}: {}, // {}".format(code, advance, unichr(code)).encode("utf8")




# for c in string.printable:
# 	face.load_char(c)
# 	print c, face.glyph.bitmap.width, 

