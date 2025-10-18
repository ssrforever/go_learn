# bufio
## 读取stdin : dup1
input := bufio.NewScanner(os.Stdin)
for input.Scan() {
    ..
}
这里的scan 是输入一行，并去除\n

## 读取file : dup2
f, err := os.Open(arg)
input := bufio.NewScanner(f)
for input.Scan() {
    ..
}


# ioutil : dup3
data, err := ioutil.ReadFile(filename)
这里读取的是整个文件内容,是一个list