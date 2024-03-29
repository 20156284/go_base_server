import Vue from 'vue'
import App from './App.vue'
// 引入element
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// 全局配置elementui的dialog不能通过点击遮罩层关闭
ElementUI.Dialog.props.closeOnClickModal.default = false
Vue.use(ElementUI);
// 引入封装的router
import router from '@/router/index'

// time line css
import '../node_modules/timeline-vuejs/dist/timeline-vuejs.css'

import '@/permission'
import { store } from '@/store/index'
Vue.config.productionTip = false

// 路由守卫
import Bus from '@/utils/bus.js'
Vue.use(Bus)

import APlayer from '@moefe/vue-aplayer';

Vue.use(APlayer, {
    defaultCover: 'https://github.com/u3u.png',
    productionTip: true,
});


import { auth } from '@/directive/auth'
// 按钮权限指令
auth(Vue)

import uploader from 'vue-simple-uploader'
Vue.use(uploader)

export default new Vue({
    render: h => h(App),
    router,
    store
}).$mount('#app')

//引入echarts
import echarts from 'echarts'
Vue.prototype.$echarts = echarts;

console.log(`
       欢迎使用 GoBaseServer
	   当前版本:V0.1.0
	   联系方式：willzh@live.cn
	   默认自动化文档地址:http://127.0.0.1%s/swagger
	   默认前端文件运行地址:http://127.0.0.1:8989
`)