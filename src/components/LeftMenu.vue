<template>

  <div class="left-content">
    <!-- 添加按钮 -->
    <div class="fileCtl">
      <n-button style="width:100%" color="#8a2be2" size="medium" @click="showModal = true">
        <template #icon>
          <n-icon>
            <LinkOutline/>
          </n-icon>
        </template>
        添加连接
      </n-button>
    </div>
    <!-- 分隔线 -->
    <n-divider>
      <span style="font-size: 13px; color: #2b4b6b;user-select: none;">Server -> Service -> Methods</span>
    </n-divider>
    <!-- 方法列表 -->
    <n-spin :show="loading">
      <MethodMenu></MethodMenu>
    </n-spin>
  </div>

  <!-- 添加地址弹框 -->
  <n-modal v-model:show="showModal" preset="dialog" title="Dialog">
    <template #header>
      <div>添加连接</div>
    </template>
    <n-input style="margin-top: 1rem"
             placeholder="http[s]://地址"
             :allow-input="noSideSpace"
             v-model:value="address"
             clearable></n-input>
    <template #action>
      <n-button type="primary" @click="AddLink">添加</n-button>
    </template>
  </n-modal>

</template>

<script lang="ts" setup>
import {LinkOutline} from '@vicons/ionicons5/'
import MethodMenu from './MethodMenu.vue';
import {onMounted, ref} from 'vue'
import {AddFile, SetFile} from '../types/request'
import {useNotification} from 'naive-ui';
import useVStore from '../api/useVStore';
import {server, service} from '../store'

const store = useVStore()
const notification = useNotification()
let showModal = ref<boolean>(false)
let address = ref<string>('localhost:50051')
let loading = ref<boolean>(false)
const noSideSpace = (value: string) => !value.startsWith(' ') && !value.endsWith(' ')

const AddLink = async () => {
  showModal.value = false
  loading.value = true
  if (address.value.length < 2) return
  try {
    const {
      data: res
    } = await AddFile(address.value)
    let servers: server[] = store.getters.getMenuVal
    let da = res.data
    // 判断重复添加
    for (let i = 0; i < servers.length; i++) {
      if (servers[i].url === da.url) {
        notification['warning']({
          content: '服务重复',
          meta: "此连接的服务已存在",
          duration: 2500,
          keepAliveOnHover: true
        })
        loading.value = false
        return
      }
    }

    // 存储
    let server: server = da
    console.log('addMenuVal', server)
    store.commit('addMenuVal', server)
    setFile()
    loading.value = false
  } catch (error) {
    notification['error']({
      content: '获取错误',
      meta: '获取失败',
      duration: 2500,
      keepAliveOnHover: true
    })
    loading.value = false
  }
}


const setFile = async () => {
  // try {
  //   const {
  //     data: res
  //   } = await SetFile({data: store.getters.getMenuVal})
  //   console.log(res)
  // } catch (error) {
  //   console.log(error)
  // }
}

</script>

<style scoped>
.left-content {
  margin: 10px;
  height: 100%;
}

.n-divider--title-position-center {
  margin: 1rem 0;
}
</style>