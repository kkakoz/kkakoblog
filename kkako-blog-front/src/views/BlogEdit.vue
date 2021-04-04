<template>
  <div class="blog">
    <el-form ref="form" :model="form" label-position="left" label-width="80px">
      <el-form-item label="标题">
        <el-input maxlength="50px" v-model="form.title"></el-input>
      </el-form-item>
      <el-form-item label="分类">
        <el-select v-model="form.type" placeholder="请选择分类">
          <el-option v-for="item in types" v-bind:key="item.id" v-bind:label="item.name" v-bind:value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="标签">
        <el-checkbox-group v-model="form.tags">
          <el-checkbox v-for="item in tags" v-bind:key="item.id" v-bind:label="item.id">{{item.name}}</el-checkbox>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item>
        <no-ssr>
          <mavon-editor :ishljs="true" :toolbars="markdownOption" v-model="form.content"/>
        </no-ssr>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">立即创建</el-button>
        <el-button>取消</el-button>
      </el-form-item>

    </el-form>

  </div>
</template>

<script>
export default {
  name: "blogedit",
  data() {
    return {
      markdownOption: {
        bold: true, // 粗体
        italic: true, // 斜体
        header: true, // 标题
        underline: true, // 下划线
        strikethrough: true, // 中划线
        mark: true, // 标记
        superscript: true, // 上角标
        subscript: true, // 下角标
        quote: true, // 引用
        ol: true, // 有序列表
        ul: true, // 无序列表
        link: true, // 链接
        imagelink: true, // 图片链接
        code: true, // code
        table: true, // 表格
        fullscreen: true, // 全屏编辑
        readmodel: true, // 沉浸式阅读
        htmlcode: true, // 展示html源码
        help: true, // 帮助
        /* 1.3.5 */
        undo: true, // 上一步
        redo: true, // 下一步
        trash: true, // 清空
        save: true, // 保存（触发events中的save事件）
        /* 1.4.2 */
        navigation: true, // 导航目录
        /* 2.1.8 */
        alignleft: true, // 左对齐
        aligncenter: true, // 居中
        alignright: true, // 右对齐
        /* 2.2.1 */
        subfield: true, // 单双栏模式
        preview: true, // 预览
      },
      types: [
        {
          id:1,
          name:"分类1",
        },
        {
          id:2,
          name:"分类2",
        }
      ],
      tags: [
        {
          id:1,
          name:"tag1",
        },
        {
          id:2,
          name:"tag2",
        }
      ],
      form: {
        title: '',
        type: '',
        tags: [],
        content: "",
      }
    };
  },
  methods: {
    onSubmit() {
      console.log('submit!');
      console.log(this.form)
    }
  }
}
</script>