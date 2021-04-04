<template>
    <div class="header">
      <!-- <div id="wrapper" :style="this.sidebar?'display:block':'display:none'">
        <i class="el-icon-close close" @click="closeSidebar" style="font-size:35px"></i>
        <div>
          <img>
        </div>
      </div> -->
    <!-- <el-col class="hidden-sm-and-down" :md="7" :lg="7" :xl="7">
      <el-button type="text" ></el-button>
    </el-col> -->
    
    <el-drawer
      title="kkako!"
      :visible.sync="drawer"
      direction="ltr"
      size="30%">
      <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose">
        <el-menu-item index="1">
          <i class="el-icon-menu"></i>
          <span slot="title">首页</span>
        </el-menu-item>
        <el-submenu index="2">
          <template slot="title">
            <i class="el-icon-location"></i>
            <span>分类</span>
          </template>
            <el-menu-item index="1-1">选项1</el-menu-item>
            <el-menu-item index="1-2">选项2</el-menu-item>
            <el-menu-item index="1-3">选项3</el-menu-item>
            <el-menu-item index="1-4-1">选项1</el-menu-item>
        </el-submenu>
        <el-menu-item index="2">
          <i class="el-icon-menu"></i>
          <span slot="title">关于我</span>
        </el-menu-item>
        <el-menu-item index="3">
          <i class="el-icon-document"></i>
          <span slot="title">登录</span>
        </el-menu-item>
      </el-menu>
    </el-drawer>
      <div class="leftbox">
        <el-col class="hidden-md-and-up" :xs="24" :sm="24">
          <div id=title>
            <i class="el-icon-s-unfold" @click="drawer = true"></i>
          </div>
        </el-col>
        <el-col class="hidden-sm-and-down" :md="24" :lg="24" :xl="24">
          <div id="title" @click="home">kkako</div>
        </el-col>
      </div>
      <div class="midbox">
        <img id="avatar" src="@/assets/avatar.webp">
      </div>
      <div class="rightbox">
        <el-col class="hidden-sm-and-down flex-end" :md="24" :lg="24" :xl="24">
          <el-menu 
            :default-active="activeIndex"
            class="el-menu-demo"
            mode="horizontal"
            @select="handleSelect"
          >
            <el-menu-item index="1">首页</el-menu-item>
            <el-menu-item index="2">
              <el-dropdown trigger="click" @command="handleCommand">
                <span class="el-dropdown-link">
                  分类<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item v-for="item in categorys" v-bind:key="item.id" command="blogEdit">写博客</el-dropdown-item>
                  <el-dropdown-item command="logout">退出登录</el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </el-menu-item>
            <el-menu-item index="3">时间轴</el-menu-item>
            <el-menu-item index="4">关于我</el-menu-item>
            <el-menu-item index="5"><i class="el-icon-search"></i></el-menu-item>
            <el-menu-item v-if="!$store.state.user.username" index="6">
              登录
            </el-menu-item> 
            <el-menu-item v-else index="7">
              <el-dropdown trigger="click" @command="handleCommand">
                <span class="el-dropdown-link">
                  <i class="el-icon-s-operation"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item command="blogEdit">写博客</el-dropdown-item>
                  <el-dropdown-item command="logout">退出登录</el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </el-menu-item>

          </el-menu>
        </el-col>

          
      </div>

      <el-dialog 
      :visible.sync="dialogFormVisible"
      >
        <el-form :model="form">
          <el-form-item label="邮箱">
            <el-input v-model="form.email"></el-input>
          </el-form-item>
          <el-form-item label="密码" >
            <el-input type="password" v-model="form.pwd" autocomplete="off"></el-input>
          </el-form-item>
        </el-form>
        <!-- <a :href="'https://github.com/login/oauth/authorize?client_id=' + githubClientId + '&redirect_uri=' + githubRedirectUrl" target="_blank"> -->
        <a @click="githubLogin" href="#">
          <svg
            style="width: 1.5em!important;height: 1.5em!important;"
            t="1588657335874"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="6878"
            width="200"
            height="200"
          >
            <path
              d="M0 520.886c0-69.368 13.51-135.697 40.498-199.02 26.987-63.323 63.322-117.826 109.006-163.51 45.65-45.65 100.154-81.985 163.51-109.006A502.289 502.289 0 0 1 512 8.92c69.335 0 135.663 13.477 198.986 40.497 63.356 26.988 117.86 63.323 163.51 109.007 45.684 45.65 82.02 100.154 109.006 163.51A502.289 502.289 0 0 1 1024 520.852c0 111.318-32.504 211.472-97.511 300.494-64.975 88.989-148.48 150.825-250.484 185.476-5.351 0-9.348-0.99-11.99-2.973-2.676-1.982-4.196-3.997-4.526-6.012a59.458 59.458 0 0 1-0.495-8.984 7.663 7.663 0 0 1-0.991-3.006v-128.99c0-40.63-14.336-75.314-43.008-103.986 76.667-13.345 134.011-41.819 171.999-85.487 37.987-43.669 57.013-96.52 57.013-158.522 0-58.005-18.663-108.346-56.022-150.99 13.345-42.678 11-87.668-6.97-135.003-18.697-1.322-39.011 1.85-61.01 9.513-22 7.663-38.318 14.831-49.02 21.47-10.637 6.673-20.316 13.016-28.97 19.027-38.68-10.669-81.854-16.02-129.486-16.02-47.7 0-90.509 5.351-128.529 16.02-7.333-5.35-15.855-11.164-25.5-17.507-9.68-6.342-26.493-14.005-50.507-22.99-23.982-9.018-45.65-12.85-65.008-11.495-18.663 47.996-20.645 93.646-5.979 136.984-36.665 42.678-54.998 92.986-54.998 150.99 0 62.002 18.663 114.689 55.99 157.994 37.326 43.339 94.67 72.01 171.998 86.016a142.303 142.303 0 0 0-39.969 70.029c-56.683 13.972-96.355 3.963-119.015-30.06-42.017-61.308-79.674-83.307-113.003-65.965-4.69 4.657-3.997 9.48 1.982 14.501 6.012 4.988 14.996 11.66 27.02 19.985 11.99 8.357 20.976 17.507 26.987 27.515 0.661 1.322 2.51 6.177 5.517 14.502a831.917 831.917 0 0 0 8.985 23.981c2.973 7.663 8.654 16.186 17.011 25.5 8.324 9.349 18.003 17.178 29.003 23.52 11 6.309 26.161 11 45.485 14.006 19.324 2.972 41.323 3.138 65.998 0.495v100.484c0 0.991-0.165 2.643-0.495 5.021-0.33 2.312-0.991 3.964-1.982 4.955-0.991 1.024-2.345 2.015-4.03 3.039a12.52 12.52 0 0 1-6.474 1.486c-2.676 0-6.012-0.33-10.009-0.99-101.343-35.345-183.825-97.182-247.51-185.51C31.842 731.037 0 631.577 0 520.92z"
              p-id="6879"
            />
          </svg>
        </a>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="login">登 录</el-button>
        </div>
      </el-dialog>
    </div>
</template>

<script>
import {BlogLogin, GithubConf, OAuthLogin} from '../api/request'

export default {
  name: 'NavBar',
  data () {
    return {
      activeIndex: '1',
      dialogFormVisible: false,
      form: {
        email: '',
        pwd: '',
      },
      categorys: [],
      drawer: false,
      githubClientId: "",
      githubRedirectUrl: "",
    }
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(this.$store.state.user)
      if (key == 1) {
        this.$router.push('/')
      }
      if (key == 3) {
        this.$router.push('/time_line')
      }
      if (key == 5) {
        this.$router.push('/blog')
      }
      if (key == 6) {
        this.dialogFormVisible = true
      }
    },
    home() {
      this.$router.push('/')
    },
    handleClose() {
      console.log('close')
    },
    login() {
      BlogLogin(this.form).then(res => {
        if (res.data.code == 200) {
          this.$message({
            showClose: true,
            message: '登录成功',
            type: 'success'
          });
          localStorage.setItem('user', res.data.data)
          this.$store.commit('login', res.data.data)
          this.dialogFormVisible = false
        } else {
          this.$message.error(res.data.msg);
        }        
      })
    },
    handleCommand(command) {
      if (command == 'blogEdit') {
        this.$router.push('/blog_edit')
      } else if (command == 'logout') {
        this.$store.commit('logout')
      }
    },
    closeSidebar() {
      this.sidebar=false
    },
    githubLogin() {
      var _this = this
      var url = 'https://github.com/login/oauth/authorize?client_id=' + _this.githubClientId//+ '&redirect_uri=' + _this.githubRedirectUrl + '/'
      window.location.href = url
    },
    getParameter() {
      var href = window.location.href.split('&')
      if (href.length == 2) {
        var querys = href[1].split('?')
        var typekey = 'auth_type'
        var codekey = 'code'
        var params = {}
        querys.forEach((query) => {
          var param = query.split('?')
          param.forEach((p) => {
            var keyvalue = p.split('=')
            if (keyvalue[0] == typekey) {
              params.auth_type = keyvalue[1]
            }
            if (keyvalue[0] == codekey) {
              params.code = keyvalue[1]
            }
          })
        })
        if (params.auth_type != '' && params.code != '') {
          return params
        }
      }
      return null
    }
  },
  mounted() {
    GithubConf().then(res => {
      this.githubClientId = res.data.client_id
      this.githubRedirectUrl = res.data.redirect_url
    })
    console.log(window.location.href)
    console.log("param")
    var a = this.getParameter()
    console.log(a)
    if (a != null) {
      OAuthLogin(a).then((res) => {
        if (res.data.code == 200) {
          this.$message({
            showClose: true,
            message: '登录成功',
            type: 'success'
          });
          localStorage.setItem('user', res.data.data)
          this.$store.commit('login', res.data.data)
          this.dialogFormVisible = false
          this.$router.push('/');
        } else {
          this.$message.error(res.data.msg);
        }        
      })
    }

  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.header {
  height: 70px;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-between;
  box-shadow: 0px 1px 0px 0px #e5e5e5;
  background-color: #FFF;
}

#wrapper {
    position: fixed;
    width: 25vw;
    height: 100vh;
    max-width: 400px;
    /* -webkit-transition: all .5s ease;
    -moz-transition: all .5s ease;
    -ms-transition: all .5s ease;
    -o-transition: all .5s ease;
    transition: all .5s ease; */
    z-index: 999;
    opacity: 0.8;
    background-color: rgba(255,255,255,.6);
}

.leftbox {
  width: 45%;
  /* display: inline-flex; */
  justify-content: flex-start;
  display: flex;
  margin-top: 10px;
}

#title {
  padding: 10px 20px;
  font-size: 30px;
  text-align: left;
  font-family: microsoft yahei;
  /* font-family: Arial, Helvetica, sans-serif; */
  color: #42b983;
}

#avatar {
  height: 70px;
  border-radius: 32.5px;
}

.rightbox {
  width: 45%;
  flex-direction: row-reverse;
  margin-top: 10px;
}

.el-menu.el-menu--horizontal{
  border-bottom: none !important;
}

.midbox {
  height: 30px;
  display: inline-flex;
}

.flex-end {
  display: flex;
  justify-content: flex-end
}
</style>
