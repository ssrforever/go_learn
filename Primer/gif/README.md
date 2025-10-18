# 创建gif
主要的使用到的库为image/gif和image/color
前者生成gif，后者设置颜色
```
anim.Delay = append(anim.Delay, delay) //设置帧间的延迟，单位是10ms
gif.EncodeAll(out, &anim) // 生成gif
```