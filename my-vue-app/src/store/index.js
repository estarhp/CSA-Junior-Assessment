import {createStore} from "vuex";
import axios from "axios";


const state = {
  isLogin:""
}

const mutations ={
  isLogin(state,message){
      state.isLogin = message
  },
    userDetails(state,details){
      state.userDetails = details
    }
}

const actions = {
   AlreadyLogin: async function(context){
      const data = await axios({
          url:"/api/isLogin"
      }).catch(err => {
          console.error(err)
          return
      })

       context.commit("isLogin",data.data.message)
       await context.dispatch("getUserDetails")

  },
  getUserDetails:async function(context){
       const result = await axios({
           url:"/api/userDetails",
           method:"get"
       })

      context.commit("userDetails",result.data.userDetails)
  }
}

export default createStore({
    state,
    actions,
    mutations
})