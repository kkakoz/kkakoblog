<template>
  <div class="time_line">
    <div class="block">
      <el-timeline :reverse="true">
        <el-timeline-item
          v-for="(activity, index) in activities"
          :key="index"
          :timestamp="activity.timestamp">
          <el-card>
            <h4>{{activity.content}}</h4>
            <p>王小虎 提交于 {{activity.timestamp}}</p>
          </el-card>
        </el-timeline-item>
      </el-timeline>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        
        activities: [{
          content: '活动按期开始',
          timestamp: '2018-04-15 10:12:15'
        }, {
          content: '通过审核',
          timestamp: '2018-04-15 10:12:15'
        }, {
          content: '创建成功',
          timestamp: '2018-04-15 10:12:15'
        },
        {
          content: '活动按期开始',
          timestamp: '2018-04-15 10:12:15'
        }, {
          content: '通过审核',
          timestamp: '2018-04-15 10:12:15'
        }, {
          content: '创建成功',
          timestamp: '2018-04-15 10:12:15'
        },
        {
          content: '活动按期开始',
          timestamp: '2018-04-15 10:12:15'
        }, {
          content: '通过审核',
          timestamp: '2018-04-15 10:12:15'
        }, {
          content: '创建成功',
          timestamp: '2018-04-15 10:12:15'
        }
        ]
      };
    }
  };
</script>

<style scoped>
.block {
  margin-top: 20px;
  
}
</style>