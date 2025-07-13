App.Vue是主体部分

## Vue主要特点
- 轻量级框架，提供MVVM数据绑定和可组合组建系统，有简单灵活的API
- 双向数据绑定，可以用简洁的模版声明渲染api
- 组件化（component），可以组成父子通信等等功能
- 客户端路由Vue Router 使路径和组件映射

[Vue的intro，官网](https://vuejs.org/guide/introduction.html)
[中文版intro](https://cn.vuejs.org/guide/introduction.html)

## Vue相关基本概念

### Single-File Components（SFC）单文件组件模式
理解成就是html，css，js都被整合到.vue文件中
#### API style
分为option ap 选项式 APIi和 composition api 组合式 API


较为重要的概念包括：
- 组件（Component）
- ​响应式数据绑定（Reactivity）​
	- 包括v-model，computed，watch等
- 指令系统（Directives）
	- 以v-为前缀等特殊属性，用于扩展html功能包括v-if，v-for，v-on等
- ​​插槽（Slots）​ 
	- 允许父组件向子组件注入内容，实现灵活的组合模式
- ​​生命周期钩子（Lifecycle Hooks）​
	- 包括mounted，created，updated等状态，是控制组件从创建到销毁的各个阶段逻辑
- 状态管理（State Management）
	- 包括state，mutation，action等
- 路由（Routing）
	- 实现页面切换，单页应用（SPA）的页面切换和导航守卫等
- ​​虚拟DOM（Virtual DOM）​
	- 优化渲染性能，通过Diff算法最小化真实DOM实现
	- 特征：render()
- ​​Composition API（Vue 3）
	- 用于地带Options API的组织方式
- 过渡与动画（Transitions）​
	- 通过\<transition>组件实现元素进入/离开的动画效果
 

## Vue安装配置
Vue安装配置比较基础，个人基于mac
首先需要安装npm （node.js相关的）之前我装过了，搜一下一个指令就行
 `npm install -g @vue/cli` 安装Vue的cli
``` bash
vue create todo-app # 在指定的目录下
cd todo-app 
npm run serve
```
这是一种面对稳定的vue2的做法

```bash
# 1. 初始化项目（回答交互问题选择需要的功能）
npm init vue@latest my-vue-app
# 2. 进入项目并安装依赖
cd my-vue-app
npm install
# 3. 启动开发服务器
npm run dev
```
这是官方面对vue3的新建方法

我感觉采取vue2的做法更加简便，于是用第一种创建方法
同时先采取组合式 API + SFC方式，就是：单文件组件（不选html）和组合式（不选选项式）书写文件
特点是：\<script setup> 以及文件为App.vue

创建对应的文件夹后，修改


目前较新的cdn：（直接官网的quickstart看就可以的）
``` html
<script src="https://unpkg.com/vue@3/dist/vue.global.js"></script>
```
## 补充技巧
###### html代码格式化
mac中shift+option+F可以让html代码格式化

## 基础知识
### 声明式渲染
``` vue
<script setup>
import { reactive, ref } from 'vue'
//ref和reactive用于创建响应式数据

const counter = reactive({ count: 0 })
console.log(counter.count) //会输出在控制台
counter.count++

const message = ref('Hello World!')
console.log(message.value) // "Hello World!"
message.value = 'Changed'
</script>

<template>
  <h1>{{ message }}</h1>
  <p>Count is: {{ counter.count }}</p>
</template>
```

### Attribute绑定
v-bind绑定动态值，后面的是对应的id或者class
``` vue
<script setup>
import { ref } from 'vue'

const titleClass = ref('title')
</script>

<template>
  <h1 v-bind:class="titleClass">Make me red</h1> <!-- 此处添加一个动态 class 绑定 -->
</template>

<style>
.title {
  color: red;
}
</style>
```
前面的v-bind其实也可以省略，直接:class="titleClass"


### 事件监听（事件处理）
使用 `v-on` 指令监听 DOM 事件   v-on通常也会用@代替
``` vue
<script setup>
import { ref } from 'vue'

const count = ref(0)

function increment() {
  count.value++
}
</script>

<template>
  <button @click="increment">Count is: {{ count }}</button>
</template>
```
调用script中的increment函数实现点击增加计数
使用{{}}用于引入插值



### 表单绑定
到这里可以看出来其实就是类似html中的id，class，input等，不过是另外的表现形式
下面实现一个输出随输入变化的绑定：
``` vue
<script setup>
import { ref } from 'vue'

const text = ref('')

function onInput(e) {
  text.value = e.target.value
}
</script>

<template>
  <input :value="text" @input="onInput" placeholder="Type here">
  <p>{{ text }}</p>
</template>
```

`e.target` 指向 `<input>`（触发源，除了输入还有比如按钮） 会精确获取触发事件的元素
比如即使：
``` html
<div @input="onInput">
  <input :value="text">
</div>
```
`e.currentTarget` 指向 `<div>`（绑定事件的对象） 否则target还是input

此外还可以使用v-model
``` vue
<script setup>
import { ref } from 'vue'

const text = ref('')
</script>

<template>
  <input v-model="text" placeholder="Type here">
  <p>{{ text }}</p>
</template>
```
会直接将input和test的值绑定

### 条件渲染
可以通过v-if指令有条件渲染，通过点击按钮实现切换
``` vue
<script setup>
import { ref } from 'vue'

const awesome = ref(true)

function toggle() {
  awesome.value = !awesome.value
}
</script>

<template>
  <button @click="toggle">Toggle</button>
  <h1 v-if="awesome">Vue is awesome!</h1>
  <h1 v-else>Oh no </h1>
</template>
```
### 列表渲染
使用 `v-for`渲染基于源数组的列表，就是列表中很多，都一个个展示出来
下面是一个增删任务列表的实现
``` vue
<script setup>
import { ref } from 'vue'

// 给每个 todo 对象一个唯一的 id
let id = 0

const newTodo = ref('')
const todos = ref([
  { id: id++, text: 'Learn HTML' },
  { id: id++, text: 'Learn JavaScript' },
  { id: id++, text: 'Learn Vue' }
])

function addTodo() {
  todos.value.push({ id: id++, text: newTodo.value })
  newTodo.value = ''
}

function removeTodo(todo) {
  todos.value = todos.value.filter((t) => t !== todo)
}
</script>

<template>
  <form @submit.prevent="addTodo">
    <input v-model="newTodo" required placeholder="new todo">
    <button>Add Todo</button>
  </form>
  <ul>
    <li v-for="todo in todos" :key="todo.id">
      {{ todo.text }}
      <button @click="removeTodo(todo)">X</button>
    </li>
  </ul>
</template>
```
- @submit.prevent是阻止表单默认提交，仅执行对应的方法
- .filter() :JavaScript 数组方法，返回一个新数组，包含通过回调函数测试的元素
- 这里过滤掉与 `todo` 不匹配的项。这里就是删除对应的todo元素key 是vue底层回调用的索引，虽然删掉也ok，不过这里留着好优化


### 计算属性
理解：Vue中的​​计算属性​​是​​基于依赖数据动态计算并缓存结果的属性​​，用于简化模板中的复杂逻辑（如反转字符串、计算总价）并提升性能
（一些函数可以实现的功能通过计算属性来实现是因为可以在缓存中运算，所以可以提升效率）


- 实现一个可以隐藏和显示的任务列表，同时完善checkbox：
``` vue
<script setup>
import { ref, computed } from 'vue'

let id = 0

const newTodo = ref('')
const hideCompleted = ref(false)
const todos = ref([
  { id: id++, text: 'Learn HTML', done: true },
  { id: id++, text: 'Learn JavaScript', done: true },
  { id: id++, text: 'Learn Vue', done: false }
])

const filteredTodos = computed(() => {
  return hideCompleted.value
    ? todos.value.filter((t) => !t.done)  //会进行清理
    : todos.value   //返回所有的内容
})

function addTodo() {
  todos.value.push({ id: id++, text: newTodo.value, done: false })
  newTodo.value = ''
}

function removeTodo(todo) {
  todos.value = todos.value.filter((t) => t !== todo)
}
</script>

<template>
  <form @submit.prevent="addTodo">
    <input v-model="newTodo" required placeholder="new todo">
    <button>Add Todo</button>
  </form>
  <ul>
    <li v-for="todo in filteredTodos" :key="todo.id">
      <input type="checkbox" v-model="todo.done">
      <span :class="{ done: todo.done }">{{ todo.text }}</span>
      <button @click="removeTodo(todo)">X</button>
    </li>
  </ul>
  <button @click="hideCompleted = !hideCompleted">
    {{ hideCompleted ? 'Show all' : 'Hide completed' }}
  </button>
</template>

<style>
.done {
  text-decoration: line-through;
}
</style>
```

- `todo.done` 为 `true` 时，Vue 会为 `<span>` 添加 `done` 类；为 `false` 时则移除该类 。这个done类就是在最下面的style中（css的设置），会添加有一条横线的效果
- 条件 ? 表达式1 : 表达式2   若条件为true就返回表达式1，不然就是表达式2

### 生命周期和模版引用
示范代码：
``` vue
<script setup>
import { ref, onMounted } from 'vue'

const pElementRef = ref(null)

onMounted(() => {
  pElementRef.value.textContent = 'mounted!'
})
</script>

<template>
  <p ref="pElementRef">Hello</p>
</template>
```

1. ​**​初始化阶段​**​：  
    `ref(null)` 创建一个空的响应式引用，Vue 会追踪其变化。
2. ​**​模板绑定​**​：  
    `<p ref="pElementRef">` 在渲染时将 DOM 元素赋值给 `pElementRef.value`。  （所以如果提前ref("mounted")就会被覆盖掉的）
3. ​**​挂载后操作​**​：  
    `onMounted` 确保 DOM 已存在，此时才能安全地通过 `.value` 修改元素属性。


### 侦听器

https://jsonplaceholder.typicode.com/ 是一个可以提供json返回格式等的网站，可以用于免费的测试

详细解答如下问题：
- ref(1) 和 let todoId = 1 的区别​
	ref(1)是响应式数据，存储在todoId.value中，当todoId.value 变化的时候，会自动更新DOM重新渲染。而let todoId = 1 只是普通的js变量，非响应式，修改不会触发重新渲染
- res = await fetch 和 todoData.value = await res.json()​
	`fetch` 发起网络请求，返回一个promise（表示异步操作），而await会等待过程完成然后再赋值给res。对于await res.json()道理也是一样的
- 单独调用 fetchData()​ 
	在 \<script setup> 中直接调用 fetchData()，相当于在组件​​初始化时自动执行​​（类似 created 生命周期）。 在这里会直接获取todoId =1 的数据。
- :disabled="!todoData"
	当`todoData.value` 为 `null`（初始状态或加载中），按钮禁止使用，防止用户在数据加载时重复点击
- \<p> 和\<pre>
	p代表普通段落，文本会​​合并空格和换行​​，适合显示简单文字，而pre会保留文本的原始格式


### 混入
Vue 的混入（mixin）是一种分发 Vue 组件中可复用功能的方式。通过混入，一个混入对象可以包含多个组件共用的选项，如数据、方法、生命周期钩子等，当组件使用这个混入时，会将混入对象的选项与自身的选项进行合并。

打个比方：类似共享背包：混入如同一个共享的背包，多个旅行者（组件）都可以背这个背包，从而拥有背包里的物品（混入的属性和方法），但旅行者自己也可以带一些独特的物品（组件自身的选项覆盖混入的选项）。



## 高级知识

### Vue Router
每个项目中都需要安装 ：  `npm install vue-router@next --save`
vue3对应4.x版本，vue2对应3.x版本
如果用CDN方式引入：
``` vue
<!-- Vue 3 + Vue Router 4   用这个就可以的  --> 
<script src="https://cdn.jsdelivr.net/npm/vue@3/dist/vue.global.js"></script>
<script src="https://cdn.jsdelivr.net/npm/vue-router@4/dist/vue-router.global.js"></script>

<!-- Vue 2 + Vue Router 3 -->
<script src="https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.js"></script>
<script src="https://cdn.jsdelivr.net/npm/vue-router@3.5.1/dist/vue-router.js"></script>
```



### Axios实现ajax
axois结构体发送不同的请求


### Vue cli 脚手架创建项目
理解文件夹结构，导入组件，静态资源等



### 状态管理
使用vuex
涉及到大型项目的数据流动，有比如commit，state，mutation，Actions等概念，中小型项目中利用emit和props其实就可以解决
暂时先没了解

#### build a dist file
use nginx
``` bash
npm install
npm run build
```
upload the dist file 