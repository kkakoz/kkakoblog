<template>
  <div class="home">
    <el-row type="flex" justify="space-between">
      <el-col :xs="24" :sm="24" :md="16" :lg="16" :xl="16">
        <div class="articles">
          <div v-for="(item, index) in articles" v-bind:key="item.id" 
          :class="index%2!=0?'article row-recover':'article row'"
          >
            <div class="leftbox">

              <img src="../assets/avatar.webp" height="100%">
            </div>
            
            <div class="rightbox">
              <div class="timeh">
                <i class="el-icon-time"></i>&nbsp;&nbsp;<p class="inlineflex">{{item.create_time}}</p>
              </div>
              <div class="title"><b>{{item.title}}</b></div>
              <div class="type"><i class="el-icon-folder"></i>&nbsp;<span style="color: #858585;">我的分类</span></div>
              <div class="content">
                {{item.content}}
              </div>
            </div>
          </div>
        </div>
      </el-col>
      <el-col class="hidden-sm-and-down" :md="7" :lg="7" :xl="7">
        <right-menu></right-menu>
      </el-col>
    </el-row>
    <!-- :class="$store.state.width>900?'home':''"> -->
    

  </div>
</template>

<script>
import RightMenu from '../components/RightMenu.vue'
// @ is an alias to /src

export default {
  name: 'Home',
  components: {
    RightMenu

  },
  data() {
    return {
      articles: [
        {
          id: 1,
          title: "我的第一篇博客",
          type: "分类",
          content: "内容",
          create_time: "2020-1-5 11:01",
        },
        {
          id: 2,
          title: "我的第一篇博客",
          type: "分类",
          content: "内容",
          create_time: "2020-1-5 11:01",
        },
        {
          id: 3,
          title: "我的第一篇博客",
          type: "分类",
          content: "内容",
          create_time: "2020-1-5 11:01",
        },
      ]
    }
  }
}

</script>


<style scoped>

.home {
  display: flex;
  justify-content: space-between;
  font-family: 'Source Han Serif SC','Source Han Serif','source-han-serif-sc','PT Serif','SongTi SC','MicroSoft Yahei',Georgia,serif;
}

.home-recover {
  display: flex;
  flex-direction: column;
}

.article {
  opacity: 1;
  padding: auto;
  width: 100%;
  height: 300px;
  position: relative;
  margin-top: 20px;
  border-radius: 10px;
  background-color: rgba(255,255,255,0);
  box-shadow: 0 1px 20px -6px rgb(0 0 0 / 50%);
  transition: box-shadow .3s ease;
  font-family: 'Source Han Serif SC','Source Han Serif','source-han-serif-sc','PT Serif','SongTi SC','MicroSoft Yahei',Georgia,serif;
}

.row {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
}

.row-recover {
  display: flex;
  flex-direction: row-reverse;
  flex-wrap: nowrap;
}

.leftbox {
  height: 100%;
  width: 45%;
}

.rightbox {
  width: 55%;
  
  flex-direction: column;
}

.leftbox img {
  width: 100%;
  height: 100%;
  transition: all .6s;
  -o-object-fit: cover;
  object-fit: cover;
}

.timeh {
  margin: 10px 10% 0 10%;
  text-align: right;
  font-size: 15px;
}

.title {
  margin-left: 20px;
  text-align: left;
  font-size: 30px;
}

.content {
  margin-left: 20px;
  margin-top: 20px;
  text-align: left;
  font-size: 15px;
}

.type {
  margin-left: 20px;
  margin-top: 20px;
  text-align: left;
}

.inlineflex {
  display: inline-flex;
}
</style>



