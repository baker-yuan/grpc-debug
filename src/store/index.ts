import {createStore} from 'vuex'


// 存储数据的类型
export interface storeType {
  jsonVals: jsonArrType
  returnVals: jsonArrType
  menuVals: menuArrType
  tabVals: string[],
  newTab: string
}

// 左侧菜单
export interface menuArrType {
  arr: serverType[]
}

// 服务
export interface serverType {
  key: string
  url: string         // url ip:port
  services: serviceType[] // 服务提供者
}

// 服务提供者
export interface serviceType {
  key: string
  serviceName: string
  methods: methodType[]
}

// 方法
export interface methodType {
  key: string
  methodName: string
  method_type: number
  inputType: string
  outputType: string
}

// 出入参
export interface jsonArrType {
  arr: jsonType[]
}

export interface jsonType {
  name: string
  data: object // 入参
}


const store = createStore<storeType>({
  state() {
    return {
      // 方法入参 jsonType
      jsonVals: {
        arr: []
      },
      // 执行结果 jsonType
      returnVals: {
        arr: []
      },
      // 服务信息 server
      menuVals: {
        arr: []
      },
      // 顶部tab
      tabVals: [],
      // 当前打开的tab
      newTab: 'undefined'
    }
  },
  mutations: {
    // 添加方法入参
    addJsonVal(state, val: jsonType) {
      state.jsonVals.arr.push(val)
    },
    // 重置方法入参
    setJsonVal(state, val: jsonType) {
      for (let i = 0; i < state.jsonVals.arr.length; i++) {
        if (state.jsonVals.arr[i].name === val.name) {
          state.jsonVals.arr[i].data = val.data
        }
      }
    },
    // 添加返回值
    addreturnVal(state, val: jsonType) {
      state.jsonVals.arr.push(val)
    },
    // 重置返回值
    setreturnVal(state, val: jsonType) {
      for (let i = 0; i < state.returnVals.arr.length; i++) {
        if (state.returnVals.arr[i].name === val.name) {
          state.returnVals.arr[i].data = val.data
        }
      }
    },
    // 添加左侧菜单(服务信息)
    addMenuVal(state, val: serverType) {
      state.menuVals.arr.push(val)
    },
    // 设置左侧菜单
    setMenuVal(state, val) {
      for (let i = 0; i < state.menuVals.arr.length; i++) {
        if (state.menuVals.arr[i].key === val.key) {
          state.menuVals.arr[i].services = val.childs
        }
      }
    },
    // 删除左侧菜单(服务信息)
    deleteMenuVal(state, val) {
      state.menuVals.arr.splice(val, 1)
    },
    //
    deleteMenuChildVal(state, val) {
      state.menuVals.arr[val.index].services.splice(val.ci, 1)
    },
    //
    deleteMenuByKey(state, key) {
      for (let i = 0; i < state.menuVals.arr.length; i++) {
        if (state.menuVals.arr[i].key === key) {
          state.menuVals.arr.splice(i, 1)
        }
      }
    },
    // 添加面板
    addTabVal(state, val: string) {
      state.tabVals.push(val)
    },
    // 删除面板
    deleteTabVal(state, val) {
      state.tabVals.splice(val, 1)
    },
    // 删除Undefined面板
    setUndefined(state) {
      // 最开始默认打开Undefined面板，Undefined面板也是在第0个位置
      state.tabVals.splice(0, 1)
    },
    // 打开新节点
    setNewTab(state, val: string) {
      state.newTab = val
    },
  },
  getters: {
    // 获取所有方法入参
    getJsonVal(state) {
      return state.jsonVals.arr
    },
    // 获取所有返回值
    getReturnVal(state) {
      return state.returnVals.arr
    },
    // 获取所有服务信息
    getMenuVal(state) {
      return state.menuVals.arr
    },
    // 获取顶部tab
    getTabVal(state) {
      return state.tabVals
    },
    // 获取当前打开的tab
    getNewTab(state) {
      return state.newTab
    }
  }
})

export default store 
