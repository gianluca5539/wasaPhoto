<script>
import HeaderComponent from '../components/HeaderComponent.vue';
import HomePostComponent from '../components/HomeView/HomePostComponent.vue';
export default {
  data() {
    return {
      userid: null,
      username: null,
      feeling: null,
      bio: null,
      picture: null,

      posts: [],

      postOpen: -1,

      token: null
    };
  },
  methods: {
    getStream() {
      this.$axios
        .get('/stream', {
          headers: { Authorization: 'Bearer ' + localStorage.getItem('token_wasa_1982801') }
        })
        .then((response) => {
          let posts = response.data.posts;
          this.posts = posts;
        })
        .catch((error) => {
          if (error.response.status == 403) {
            // clear local storage entries
            localStorage.removeItem('userid_wasa_1982801');
            localStorage.removeItem('username_wasa_1982801');
            localStorage.removeItem('bio_wasa_1982801');
            localStorage.removeItem('feeling_wasa_1982801');
            localStorage.removeItem('picture_wasa_1982801');
            localStorage.removeItem('token_wasa_1982801');
            // redirect to login page
            this.$router.push('/login');
          }
          alert("Couldn't get stream. Please try again later.");
        });
    },
    updatePost(postid, key, value) {
      this.posts = this.posts.map((post) => {
        if (post.postid == postid) {
          post[key] = value;
        }
        return post;
      });
    },
    removePost(postid) {
      this.posts = this.posts.filter((post) => post.postid != postid);
      this.closePost();
    },
    openPost(id) {
      document.body.classList.add('no-scroll'); // disable scrolling on page
      this.postOpen = id;
    },
    closePost() {
      this.postOpen = -1;
      document.body.classList.remove('no-scroll'); // enable scrolling on page
    }
  },
  components: {
    HeaderComponent: HeaderComponent,
    HomePostComponent
  },
  created() {
    this.userid = parseInt(localStorage.getItem('userid_wasa_1982801'));
    this.username = localStorage.getItem('username_wasa_1982801');
    this.bio = localStorage.getItem('bio_wasa_1982801');
    this.feeling = parseInt(localStorage.getItem('feeling_wasa_1982801'));
    this.picture = parseInt(localStorage.getItem('picture_wasa_1982801'));
    this.token = localStorage.getItem('token_wasa_1982801');

    this.getStream();
  }
};
</script>

<template>
  <div class="homepage-container">
    <HeaderComponent
      :userid="this.userid"
      :username="this.username"
      :feeling="this.feeling"
      :picture="this.picture"
    />
    <div class="homepage-feed-container">
      <div class="homepage-feed-spacer"></div>
      <HomePostComponent
        v-if="!this.posts"
        :postid="0"
        :userid="0"
        :picture="-2"
        name="No posts yet!"
        :feeling="4"
        :userPicture="-2"
        :date="0"
        caption="Start following people to see their posts here!"
        :likeCount="999"
        :updatePost="() => {}"
        :removePost="() => {}"
        :openPost="() => {}"
        :closePost="() => {}"
        :postOpen="false"
      />
      <HomePostComponent
        v-for="post in this.posts"
        :key='post.postid'
        :postid="post.postid"
        :userid="post.userid"
        :picture="post.picture"
        :name="post.username"
        :feeling="post.feeling"
        :userPicture="post.userpicture"
        :date="post.createdat"
        :caption="post.caption"
        :likeCount="post.likecount"
        :updatePost="this.updatePost"
        :removePost="() => this.removePost(post.postid)"
        :openPost="() => this.openPost(post.postid)"
        :closePost="this.closePost"
        :postOpen="this.postOpen === post.postid"
      />
    </div>
  </div>
</template>

<style lang="scss">
body.no-scroll {
  overflow: hidden;
}
.homepage-container {
  width: 100vw;
  max-width: 100vw;
  display: flex;
  flex-direction: column;
  background: url('../assets/images/homepage-background.jpg') no-repeat center
    center fixed;
  .homepage-feed-container {
    margin-top: 20px;
    display: flex;
    flex-direction: column;
    width: 100%;
    min-height: 100vh;
    height: fit-content;
    align-items: center;
    .homepage-feed-spacer {
      width: 100%;
      height: 0px;
      transition: height 0.3s ease;
      @media screen and (max-width: 1250px) {
        height: 30px;
      }
    }
  }
}
</style>
