# go-movie

## create sample

##### create sample.png

```
convert -size 640x854 -background "#C0C0C0" -fill "#FFFF00" caption:"test" ./sample.png
```

##### create sample.mp4

```
ffmpeg -stream_loop -1 -i ./sample.png -t 30 sample.mp4
```

##### create sample.m3u8

```
ffmpeg -i sample.mp4 -f segment -segment_time 10 -segment_list sample.m3u8 -segment_format mpegts sample%03d.mp4
```
