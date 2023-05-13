<template>
  <!-- 搜索 -->
  <n-input-group>
    <n-input v-model:value="pattern" placeholder="搜索" size="small" style="margin-bottom: 1rem" />
  </n-input-group>
  <!-- 方法列表
  https://www.naiveui.com/zh-CN/light/components/tree
  -->
  <n-tree :show-irrelevant-nodes="true"
          :pattern="pattern"
          :data="data.data"
          block-line
          v-if="bShow"
          :node-props="ClickTree" />
  <!-- 无数据展示 -->
  <n-empty description="无数据" v-if="!bShow">
    <template #icon>
      <n-icon>
        <ReceiptOutline />
      </n-icon>
    </template>
  </n-empty>
  <!-- 右键菜单 -->
  <n-dropdown trigger="manual"
              placement="bottom-start"
              :show="showDropdown"
              :options="(options as any)" :x="x" :y="y"
              @select="HandleSelect"
              @clickoutside="handleClickOutside" />
</template>

<script setup lang="ts">
import { h, ref, reactive, watch, onMounted } from 'vue'
import { DropdownOption, NIcon, TreeOption, useNotification } from 'naive-ui'
import { InformationCircleOutline, ReceiptOutline, Infinite, List, GitNetwork, ListOutline } from '@vicons/ionicons5'
import TreeNode from '../types/treenodes'
import useVStore from '../api/useVStore'
import { AddFile, SetFile, GetMethodParam } from '../types/request'
import { server, service, jsonType } from '../store'

const notification = useNotification()
const store = useVStore()
const data: TreeNode = reactive(new TreeNode)
let bShow = ref(false)
let pattern = ref('')
let showDropdown = ref(false)
let x = ref(0)
let y = ref(0)
let options = ref<DropdownOption[]>([])
let clickOption = reactive<TreeOption>({})
options.value = [
  { key: 'delete', label: '删除' },
  { key: 'update', label: '刷新' },
  { key: 'copy', label: '复制' }
]

const handleClickOutside = () => {
  showDropdown.value = false
}

// 添加树形数据
// https://www.naiveui.com/zh-CN/light/components/tree
// https://www.naiveui.com/zh-CN/light/components/icon
const addMenu = (server) => {
  // 服务
  let grandpa: TreeOption = reactive({
    label: server.url,
    key: server.key,
    prefix: () => h(NIcon, { size: '1.2rem', color: '#18A058' }, () => h(ListOutline)),
    children: []
  })
  // 服务提供者
  for (let i = 0; i < server.services.length; i++) {
    let service = server.services[i]
    let father = {
      nodeType: 'service',
      label: service.serviceName,
      key: service.key,
      prefix: () => h(NIcon, { size: '1.2rem', color: '#b2de27' }, () => h(InformationCircleOutline)),
      children: []
    }
    for (let j = 0; j < service.methods.length; j++) {
      let method = service.methods[j]
      let grandson = {
        nodeType: 'method',
        label: method.methodName,
        key: method.key,
        prefix: () => h(NIcon, { size: '1.2rem', color: '#b2de27' }, () => h(Infinite))
      }
      father.children.push(grandson)
    }
    grandpa.children.push(father)
  }
  data.PushOption(grandpa)
}

/**
 * 添加服务
 *
 * @param server server
 * @constructor
 */
const SetMenu = (server) => {
  for (let i = 0; i < server.length; i++) {
    addMenu(server[i])
  }
}

// 监听
watch(store.getters.getMenuVal, (newVal, oldVal) => {
  data.Clear()
  SetMenu(newVal)
}, { immediate: true, deep: true })

watch(data.data, (newVal, oldVal) => {
  if (newVal.length > 0) {
    bShow.value = true
  } else {
    bShow.value = false
  }
})

onMounted(() => {
  if (store.getters.getMenuVal.length > 0 && data.data.length === 0) {
    data.Clear()
    SetMenu(store.getters.getMenuVal)
    bShow.value = true
  } else if (store.getters.getMenuVal.length > 0 && data.data.length > 0) {
    bShow.value = true
  }
})

// 保存数据
const setFile = async () => {
  // try {
  //   const {
  //     data: res
  //   } = await SetFile({ data: store.getters.getMenuVal })
  //   console.log(res)
  // } catch (error) {
  //   console.log(error)
  // }
}

/**
 * 获取方法入参
 *
 * @param serviceName 服务名
 * @param methodName 方法名
 * @param url 地址 ip:port
 * @param key key
 * @constructor
 */
const GetParam = async (serviceName, methodName, url, key) => {
  try {
    const {
      data: res
    } = await GetMethodParam(serviceName, methodName, url)
    let newVal: jsonType = {
      name: key,
      data: JSON.parse(res.data)
    }
    store.commit('setJsonVal', newVal)
  } catch (error) {
    notification['error']({
      content: '错误',
      meta: '刷新错误',
      duration: 2500,
      keepAliveOnHover: true
    })
  }
}
/**
 * 货物服务信息
 *
 * @param url ip:port
 * @constructor
 */
const GetMethods = async (url) => {
  try {
    const {
      data: res
    } = await AddFile(url)
    let da = res.data
    store.commit('setMenuVal', da)
    setFile()
  } catch (error) {
    notification['error']({
      content: '错误',
      meta: '刷新错误',
      duration: 2500,
      keepAliveOnHover: true
    })
  }
}

/**
 *
 *
 * @param key
 * @param option
 * @constructor
 */
const HandleSelect = (key: string | number, option: DropdownOption) => {
  console.log('handleSelect', key, option)
  showDropdown.value = false
  let st = store.getters.getMenuVal
  if (key === 'delete') {
    if (clickOption.children) {
      store.commit('deleteMenuByKey', clickOption.key)
      setFile()
    } else {
      for (let i = 0; i < st.length; i++) {
        for (let j = 0; j < st[i].childs.length; j++) {
          let name = st[i].key + ":" + st[i].childs[j].name
          if (name === clickOption.key) {
            store.commit('deleteMenuChildVal', { index: i, ci: j })
            if (store.getters.getMenuVal[i].childs.length === 0) {
              store.commit('deleteMenuVal', i)
            }
            setFile()
            return
          }
        }
      }
    }
  } else if (key === 'copy') {
    const input = document.createElement("input");
    for (let i = 0; i < st.length; i++) {
      for (let j = 0; j < st[i].childs.length; j++) {
        let name = st[i].key + ":" + st[i].childs[j].name
        if (name === clickOption.key) {
          input.value = st[i].childs[j].name
          break
        }
      }
    }
    document.body.appendChild(input)
    input.select()
    document.execCommand("Copy")
    document.body.removeChild(input)
    notification['success']({
      content: '复制成功',
      meta: "已成功将方法名称置入剪切板",
      duration: 2500,
      keepAliveOnHover: true
    })
  } else if (key === 'update') {
    if (clickOption.children) {
      let pr = clickOption.key.toString().split("::")
      GetMethods(pr[0])
    } else {
      let me = store.getters.getMenuVal
      for (let i = 0; i < me.length; i++) {
        for (let j = 0; j < me[i].childs.length; j++) {
          let keyname = me[i].key + ":" + me[i].childs[j].name
          if (clickOption.key === keyname) {
            GetParam(me[i].childs[j].father, me[i].childs[j].name, me[i].childs[j].url, keyname)
            return
          }
        }
      }
    }
  }
}

/**
 * 节点点击事件
 *
 * @param option
 * @constructor
 */
const ClickTree = ({ option }: { option: TreeOption }) => {
  let nodeType = Reflect.get(option, 'nodeType')
  console.log('clickTree', nodeType)
  if (!(nodeType === 'method')) {
    return
  }

  return {
    onClick() {
      if (option.children) return
      store.commit('setNewTab', option.key)
    },
    onContextmenu(e: MouseEvent): void {
      clickOption = option
      console.log(option)
      showDropdown.value = true
      x.value = e.clientX
      y.value = e.clientY
      e.preventDefault()
    }
  }
}

</script>