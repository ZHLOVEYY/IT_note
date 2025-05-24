###### label for 的用法
```html
<!-- 点击标签即可聚焦输入框 -->
<label for="search">搜索：</label>
<input id="search" type="text">
```

###### target的指定
\_self：在当前窗口中打开链接（默认值）。
\_blank：在新窗口或新标签页中打开链接。
\_parent：在父框架中打开链接。
\_top：在整个窗口中打开链接，会清除所有框架。

理解self和top的区别，比如有inframe的话（内置框架），self就会在内置框架中打开，top会整个窗口重新渲染