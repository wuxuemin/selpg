# selpg
* 测试的文件为file2:共有77行

* 测试1：
输入的命令行`./selpg -s -1 -e -1 file2`

![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-1.png?raw=true)


* 测试2：
输入的命令行`./selpg -s 1 -e 1 file2`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-2.png?raw=true)


* 测试3：
输入的命令行`./selpg -s 1 -e 1 < file2`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-3.png?raw=true)

![从第一行至72行(中间省略)](https://github.com/wuxuemin/selpg/blob/master/image/1-3-2.png?raw=true)


* 测试4：
输入的命令行`./selpg -s 1 -e 1 file2 > fileout`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-4.png?raw=true)

* 测试5：
输入的命令行`./selpg -s 4 -e 3 file2 2>erroroutput`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-5-1.png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-5-2.png?raw=true)

* 测试6：
输入的命令行`./selpg -s 1 -e 3 file2 >fileout 2>erroroutput`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-6-1.png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-6-2.png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-6-3.png?raw=true)

* 测试7：
输入的命令行`./selpg -s 1 -e 2 -f file2`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-7-1.png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-7-2.png?raw=true)

* 测试8：
输入的命令行`./selpg -s 1 -e 2 file2 2>/dev/null > fileout`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/xin1-9-1(7).png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-9-2(7).png?raw=true)

* 测试9：
输入的命令行`./selpg -s 1 -e 3 file2 >/dev/null`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-10(8).png?raw=true)

* 测试10：
输入的命令行`./selpg -s 1 -e 1 file2 | wc`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-11.png?raw=true)

* 测试11：
输入的命令行`./selpg -s 1 -e 3 file2 2>error | wc`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-12-1.png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-12-2.png?raw=true)

* 测试12：
输入的命令行`./selpg -s 1 -e 1 -l 10 file2`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/2-1.png?raw=true)


* 测试13：
输入的命令行`./selpg -s 1 -e 1 file2 > fileout 2>errorfile &`
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-14.png?raw=true)

* 测试14：
输入的命令行`./selpg -s 1 -e 1 lp -d xxx`
未做打印测试
