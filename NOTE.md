```js
ffmpegPath +' -ss ' + options.time + ' -i "' + options.input+ '"' + options.size + ' -vframes 1 -f image2 "' + options.output+'"'
```