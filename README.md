# go-movie

```
curl localhost:8080
```

```
convert -size 640x854 -background "#C0C0C0" -fill "#FFFF00" caption:"test" ./sample.png
```

```
ffmpeg -stream_loop -1 -i ./sample.png -t 30 sample.mp4
```

```
ffmpeg -i sample.mp4 -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list sample.m3u8 -segment_format mpegts sample%03d.mp4
```

ffmpeg -i sample.mp4 -f segment -segment_time 10 -segment_list sample.m3u8 -segment_format mpegts sample%03d.mp4
