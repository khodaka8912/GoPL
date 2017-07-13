## パイプラインの最大段数
メモリ8GのPCで無制限にゴルーチンの作成を実行。
エラー時のダンプの最終出力は↓。
```
goroutine 3039457 [runnable]:
main.pipe(0xc684453d40, 0xc684453da0)
        C:/Users/hodaka/GoglandProjects/GoPL/ch09/ex04/routins.go:12
created by main.main
        C:/Users/hodaka/GoglandProjects/GoPL/ch09/ex04/routins.go:7 +0x95
```
約300万のゴルーチンが生成されている。

## パイプラインを伝わる時間
```
$ ./routins.exe -limit 10000
6.018100 ms has elapsed for communications of 10000 pipes
$ ./routins.exe -limit 100000
57.151900 ms has elapsed for communications of 100000 pipes
$ ./routins.exe -limit 1000000
3024.034800 ms has elapsed for communications of 1000000 pipes
```
