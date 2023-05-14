<!-- 标签tab -->
<template>
  <!--
  https://www.naiveui.com/zh-CN/light/components/tabs
  -->
  <n-tabs v-model:value="pvalue"
          type="card"
          :closable="closable"
          @close="handleClose"
          @add="handleAdd"
          @update:value="changeVal">
    <n-tab-pane v-for="panel in panels"
                :key="panel"
                :name="panel"
                display-directive="show">
      <div class="top" v-if="methodInfo">
        {{ methodInfo }}
      </div>
      <div class="content">
        <!-- 请求 -->
        <div class="left">
          <div class="qrstyle">request</div>
          <Editor :keyname="panel" type="request"></Editor>
        </div>
        <!-- 中间按钮 -->
        <n-popover trigger="hover">
          <template #trigger>
            <n-button circle secondary type="error" class="staticBtn" size="large" @click="RequestForm">
              <template #icon>
                <n-icon>
                  <CaretForward/>
                </n-icon>
              </template>
            </n-button>
          </template>
          <span>发送请求</span>
        </n-popover>
        <!-- 响应 -->
        <div class="right">
          <div class="qrstyle">response</div>
          <n-spin :show="loading">
            <Editor :keyname="panel" type="response"></Editor>
          </n-spin>
        </div>
      </div>
    </n-tab-pane>
  </n-tabs>
</template>

<script lang="ts" setup>
import {ref, reactive, computed, watch, onMounted} from 'vue'
import {CaretForward} from '@vicons/ionicons5'
import Editor from './Editor.vue'
import useVStore from '../api/useVStore';
import {GetMethodParam, Query} from '../types/request'
import {jsonType, method, server, service} from '../store'
import {useNotification} from 'naive-ui';

const notification = useNotification()
const store = useVStore()
let pvalue = ref<string>('1')
const panels = reactive<string[]>([])
let loading = ref<boolean>(false)
const closable = computed(() => {
  return panels.length > 1
})

let methodInfo = ref<{exist: boolean, server: server, service: service, method:method}>()

// 加载顶部tab栏
onMounted(() => {
  let tabs = store.getters.getTabVal
  if (panels.length === 0 && tabs.length > 0) {
    for (let i = 0; i < tabs.length; i++) {
      panels.push(tabs)
    }
    pvalue.value = tabs[0]
  }
})


// 根据id查询入参
const getBodyById = (id) => {
  let params: jsonType[] = store.getters.getJsonVal
  for (let i = 0; i < params.length; i ++) {
    if (params[i].name == id) {
      return params[i].data
    }
  }
}

// 根据id查询参数
const getCallParamById = (id) => {
  let notExist = {exist: false, server: undefined, service: undefined, method: undefined}
  if (!id) {
    return notExist;
  }
  let servers: server[] = store.getters.getMenuVal
  for (let i = 0; i < servers.length; i ++) {
    let server: server = servers[i]
    for (let j = 0; j < server.services.length; j++) {
      let service: service = server.services[j]
      for (let k = 0; k < service.methods.length; k++) {
        let method: method = service.methods[k]
        if (method.key === id) {
          return {exist: true, server: server, service: service, method: method}
        }
      }
    }
  }
  return notExist
}

// 获取方法入参
const GetParam = async (serviceName, methodName, url, key) => {
  try {
    const {
      data: res
    } = await GetMethodParam(serviceName, methodName, url)
    let newVal: jsonType = {
      name: key,
      data: JSON.parse(res.data)
    }
    store.commit('addJsonVal', newVal)
  } catch (error) {
    notification['error']({
      content: '错误',
      meta: '获取参数失败',
      duration: 2500,
      keepAliveOnHover: true
    })
  }
}

//
const handleAdd = (name: string, b: boolean) => {
  if (panels.length === 1 && panels[0] === 'undefined') {
    store.commit('setUndefined')
    panels.splice(0, 1)
  }
  panels.push(name)
  pvalue.value = name
  if (b) {
    store.commit('addTabVal', name)
  }
  let { server, service, method } = getCallParamById(name)
  if (server == undefined) {
    return
  }
  GetParam(service.serviceName, method.methodName, server.url, method.key)
}

// 处理tab关闭
const handleClose = (name: string) => {
  const nameIndex = panels.findIndex((panelName) => panelName === name)
  if (!~nameIndex) return
  panels.splice(nameIndex, 1)
  store.commit('deleteTabVal', nameIndex)
  if (name === pvalue.value) {
    pvalue.value = panels.at(-1)
  }
}

const changeVal = (name: string) => {
  const index = panels.findIndex((panelName) => panelName === name)
  pvalue.value = panels[index]
}

// 设置返回值
const setReturn = (key, res) => {
  let st = store.getters.getReturnVal
  if (st.length === 0) {
    let newVal: jsonType = {
      name: key,
      data: res
    }
    st.push(newVal)
  }
  for (let i = 0; i < st.length; i++) {
    if (key === st[i].name) {
      st[i].data = res
      loading.value = false
      return
    }
  }
  let newVal: jsonType = {
    name: key,
    data: res
  }
  st.push(newVal)
  loading.value = false
}

// 请求后端服务
const RequestForm = () => {
  let { server, service, method } = getCallParamById(pvalue.value)
  let body = getBodyById(pvalue.value)
  query(service.serviceName, method.methodName, server.url, body, pvalue.value)
}

// 请求后端服务
const query = async (serviceName, methodName, url, data, key) => {
  loading.value = true
  try {
    const {
      data: res
    } = await Query(serviceName, methodName, url, JSON.stringify(data))
    setReturn(key, JSON.parse(res.data))
  } catch (error) {
    notification['error']({
      content: '错误',
      meta: '请求失败',
      duration: 2500,
      keepAliveOnHover: true
    })
    setReturn(key, {error: error.response.data})
    loading.value = false
  }
}

// 监听tab打开事件
watch(() => store.getters.getNewTab, (newVal: string, oldVal) => {
  let v = getCallParamById(newVal)
  if (v && v.exist) {
    methodInfo.value = v;
  }


  let tabs = store.getters.getTabVal
  console.log('tabs watch', tabs, newVal)
  for (let i = 0; i < tabs.length; i++) {
    if (tabs[i] === newVal) {
      pvalue.value = newVal
      return
    }
  }
  handleAdd(newVal, true)
}, {immediate: true})

</script>

<style scoped>
.content {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  position: relative;
}

.left {
  width: 49%;
  height: calc(100vh - 98px - 2rem);
  border: 0;
  display: flex;
  flex-direction: column;
  text-align: center;
  user-select: none;
  color: #2f4f6f;
  font-size: 1rem;
}

.right {
  max-width: 49%;
  width: 49%;
  height: calc(100vh - 98px - 2rem);
  display: flex;
  flex-direction: column;
  text-align: center;
  user-select: none;
  color: #2f4f6f;
  font-size: 1rem;
}

.staticBtn {
  position: absolute;
  left: calc(50% - 20px);
  top: 50%;
  z-index: 50;
}

.qrstyle {
  border: 1px solid rgb(231, 222, 222);
}
</style>