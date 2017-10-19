# selpg
* 测试的文件为file2:共有77行

* 测试1：
输入的命令行`./selpg -s -1 -e -1 file2`

![输出](https://github.com/wuxuemin/selpg/image/1-1.png)


* 测试2：
输入的命令行`./selpg -s 1 -e 1 file2`
![图片1-2]()

* 测试3：
输入的命令行`./selpg -s 1 -e 1 < file2`
![图片1-3-1 1-3-2]()


* 测试4：
输入的命令行`./selpg -s 1 -e 1 file2 > fileout`
![图片1-4]()

* 测试5：
输入的命令行`./selpg -s 4 -e 3 file2 2>erroroutput`
![图片1-5-1 1-5-2]()

* 测试6：
输入的命令行`./selpg -s 1 -e 3 file2 >fileout 2>erroroutput`
![图片1-6-1 1-6-2 1-6-3]()

* 测试7：
输入的命令行`./selpg -s 1 -e 2 -f file2`
![图片1-7-1 1-7-2]()

* 测试8：
输入的命令行`./selpg -s 1 -e 1 -l 10 file2`
![图片2-1]()

