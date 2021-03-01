<template>
  <div id="app">
    <div v-show="loading" class="loading"></div>
    <div v-show="!loading">
      <div v-for="(tag, idx) in orders" :key="idx">
        <div class="mod">
          <h2>
            <span>{{titles[tag]}}</span> · · ·
            <span class="pl">({{datas[tag].length}}本)</span>
          </h2>
          <ul>
            <li v-for="item in datas[tag]" :key="item.id">
              <div class="book-cover">
                <img :alt="item.title" :title="item.title" :src="item.pic" referrerPolicy="no-referrer">
              </div>
              <div class="book-title">
                <a target="_blank" v-bind:href="'https://book.douban.com/subject/' + item.id">{{item.title}}</a>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
/* eslint-disable */
import axios from 'axios'
export default {
  name: "app",
  components: {},
  data() {
    return {
      datas: {
        wish: [],
        reading: [],
        read: []
      },
      titles: {
        wish: "想读",
        reading: "在读",
        read: "读过"
      },
      list: {},
      orders: ["reading", "wish", "read"],
      loading: true,
      t: null
    };
  },
  created() {
    let _this = this;
    axios
      .get("https://www.haoyizebo.com/api/douban_book.json")
      .then(response => {
        _this.datas["reading"] = response.data["do"]
        _this.datas["wish"] = response.data["wish"]
        _this.datas["read"] = response.data["collect"]
        _this.loading = false;
      });
  },
  methods: {
  }
};
</script>

<style scoped>
#app a {
  text-decoration: none;
  border-bottom: dotted 1px;
  line-height: 2;
}

#app ul {
  list-style: none;
  padding: 0px;
}

#app ul li {
  width: 135px;
  height: 250px;
  margin: 0px 20px;
  list-style: none;
  display: inline-block;
}

#app .book-cover {
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex-box;
  display: flex;
  -webkit-box-align: center;
  -ms-flex-align: center;
  align-items: center;
  -webkit-box-pack: center;
  -ms-flex-pack: center;
  justify-content: center;
  text-align: center;
  width: 135px;
  height: 200px;
  text-align: center;
  line-height: 200px;
  overflow: hidden;
}

#app .book-title {
  text-align: center;
  font-size: 14px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

#app li img {
  max-height: 100%;
  width: 100%;
}

#app .mod {
  display: block;
}

#app .loading {
  top: 50%;
  left: 50%;
  margin: 20px 50% 5px 50%;
  -webkit-transform: translateY(-50%);
  transform: translateY(-50%);
  height: 30px;
  width: 10px;
  border-radius: 10%;
  background: #2c3e50;
  border-top-color: #2980b9;
  -webkit-animation: fade2 1s infinite;
  animation: fade2 1s infinite;
  -webkit-transition: background 0.2s;
  transition: background 0.2s;
}
#app .loading:after,
#app .loading:before {
  content: "";
  position: absolute;
  display: block;
  height: 20px;
  width: 10px;
  background: #2c3e50;
  top: 70%;
  -webkit-transform: translateY(-75%);
  transform: translateY(-75%);
  left: -15px;
  border-radius: 10%;
  -webkit-animation: fade1 1s infinite;
  animation: fade1 1s infinite;
  -webkit-transition: background 0.2s;
  transition: background 0.2s;
}
#app .loading:before {
  left: 15px;
  -webkit-animation: fade3 1s infinite;
  animation: fade3 1s infinite;
}

@-webkit-keyframes fade1 {
  0% {
    background: #2980b9;
  }
}

@keyframes fade1 {
  0% {
    background: #2980b9;
  }
}
@-webkit-keyframes fade2 {
  33% {
    background: #2980b9;
  }
}
@keyframes fade2 {
  33% {
    background: #2980b9;
  }
}
@-webkit-keyframes fade3 {
  66% {
    background: #2980b9;
  }
}
@keyframes fade3 {
  66% {
    background: #2980b9;
  }
}
</style>
