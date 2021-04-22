<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="12">
        <el-card>
          <div slot="header">go_base_server</div>
         
        </el-card>
      </el-col>
      
    </el-row>
  </div>
</template>

<script>
import { Commits, Members } from "@/api/github";
import Timeline from "timeline-vuejs";
export default {
  name: "Test",
  components: {
    Timeline,
  },
  data() {
    return {
      messageWhenNoItems: "There arent commits",
      members: [],
      dataTimeline: [],
      page: 0,
    };
  },
  created() {
    this.loadCommits();
    this.loadMembers();
  },
  methods: {
    loadMore() {
      this.page++;
      this.loadCommits();
    },
    loadCommits() {
      Commits(this.page).then(({ data }) => {
        data.forEach((element) => {
          if (element.commit.message) {
            this.dataTimeline.push({
              from: new Date(element.commit.author.date),
              title: element.commit.author.name,
              showDayAndMonth: true,
              description: `<a style="color: #26191b" href="${element.html_url}">${element.commit.message}</a>`,
            });
          }
        });
      });
    },
    loadMembers() {
      Members().then(({ data }) => {
        this.members = data;
        this.members.sort();
      });
    },
  },
};
</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
    "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>
