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

      token: null
    };
  },
  methods: {
    getStream() {
      this.$axios
        .get('/stream', {
          headers: { Authorization: 'Bearer ' + localStorage.getItem('token') }
        })
        .then((response) => {
          console.log(response);
        })
        .catch((error) => {
          console.log(error);
        });
    }
  },
  components: {
    HeaderComponent: HeaderComponent,
    HomePostComponent
  },
  created() {
    this.userid = parseInt(localStorage.getItem('userid'));
    this.username = localStorage.getItem('username');
    this.bio = localStorage.getItem('bio');
    this.feeling = parseInt(localStorage.getItem('feeling'));
    this.picture = localStorage.getItem('picture');
    this.token = localStorage.getItem('token');

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
        v-for="post in posts"
        :id="post.postid"
        :picture="post.picture"
        :name="post.username"
        :feeling="post.feeling"
        :profilePicture="post.profilePicture"
        :date="post.createdat"
        :caption="post.caption"
        :likeCount="post.likecount"
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
    height: fit-content;
    align-items: center;
    .homepage-feed-spacer {
      width: 100%;
      height: 0px;
      transition: height 0.3s ease;
      @media screen and (max-width: 1150px) {
        height: 30px;
      }
    }
  }
}
</style>
