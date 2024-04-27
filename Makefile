.PHONY: run
run:
	@php -S localhost:8000

.PHONY: generate
generate:
	@convert -size 640x854 -background "#C0C0C0" -fill "#FFFF00" caption:"test" sample.png
	@ffmpeg -stream_loop -1 -i sample.png -t 30 sample.mp4
	@ffmpeg -i sample.mp4 -f segment -segment_time 10 -segment_list sample.m3u8 -segment_format mpegts sample%03d.mp4
