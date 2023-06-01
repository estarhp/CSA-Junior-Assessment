import {createStore} from "vuex";
import axios from "axios";


const state = {
  isLogin:""
}

const mutations ={
  isLogin(state,message){
      state.isLogin = message
  }
}

const actions = {
   AlreadyLogin: async function(context){
      const data = await axios({
          url:"/api/isLogin"
      })

       context.commit("isLogin",data.data.message)

  }
}

export default createStore({
    state,
    actions,
    mutations
})