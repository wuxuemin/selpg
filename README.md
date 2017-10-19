# selpg
* 测试的文件为file2:共有77行

### 测试1：
输入的命令行`./selpg -s -1 -e -1 file2`
+ 错误设置页号

![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-1.png?raw=true)


### 测试2：
输入的命令行`./selpg -s 1 -e 1 file2`
+  打印出第一页,72行
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-2.png?raw=true)


### 测试3：
输入的命令行`./selpg -s 1 -e 1 < file2`
+  打印出第一页,72行,从第一行至72行(中间省略)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-3.png?raw=true)

![从第一行至72行(中间省略)](https://github.com/wuxuemin/selpg/blob/master/image/1-3-2.png?raw=true)


### 测试4：
输入的命令行`./selpg -s 1 -e 1 file2 > fileout`
+  selpg 将第 1页写至标准输出；标准输出被 shell／内核重定向至“fileout”
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-4.png?raw=true)

### 测试5：
输入的命令行`./selpg -s 4 -e 3 file2 2>erroroutput`
+ selpg 将错误的所有内容都写至“erroroutput”
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-5-1.png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-5-2.png?raw=true)

### 测试6：
输入的命令行`./selpg -s 1 -e 3 file2 >fileout 2>erroroutput`
+ selpg 将错误的所有内容都写至“erroroutput”,标准输出写至“fileout"
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-6-1.png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-6-2.png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-6-3.png?raw=true)

### 测试7：
输入的命令行`./selpg -s 1 -e 2 -f file2`
+ 页由换页符划分。第 10页到第 20页显示在屏幕上。
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-7-1.png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-7-2.png?raw=true)

### 测试8：
输入的命令行`./selpg -s 1 -e 2 file2 2>/dev/null > fileout`
+ selpg 将错误的所有内容都写至/dev/null（空设备）丢弃,标准输出写至“fileout"
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/xin1-9-1(7).png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-9-2(7).png?raw=true)

### 测试9：
输入的命令行`./selpg -s 1 -e 3 file2 >/dev/null`
+ selpg 将错误消息显示,检查错误消息"
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-10(8).png?raw=true)

### 测试10：
输入的命令行`./selpg -s 1 -e 1 file2 | wc`
+ 显示选定范围的页中包含的行数、字数和字符数
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-11.png?raw=true)

### 测试11：
输入的命令行`./selpg -s 1 -e 3 file2 2>error | wc`
+ 与上面的测试10 相似，只有一点不同：错误消息被写至“error”中
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-12-1.png?raw=true)
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-12-2.png?raw=true)

### 测试12：
输入的命令行`./selpg -s 1 -e 1 -l 10 file2`
+ 将页长设置为 10 行，第1页被显示在屏幕上。
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/2-1.png?raw=true)


### 测试13：
输入的命令行`./selpg -s 1 -e 1 file2 > fileout 2>errorfile &`
+ 运行进程,进程标识显示
![输出为:](https://github.com/wuxuemin/selpg/blob/master/image/1-14.png?raw=true)

### 测试14：
输入的命令行`./selpg -s 1 -e 1 -d lp xxx`
+ 该命令将使输出在打印机上打印。
未做打印测试
