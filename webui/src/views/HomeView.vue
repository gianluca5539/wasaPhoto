<script>
import HeaderComponent from '../components/HeaderComponent.vue';
import HomePostComponent from '../components/HomeView/HomePostComponent.vue';
export default {
  data() {
    return {
      userid: null,
      username: null,
      feeling: null,
      picture: null
    };
  },
  components: {
    HeaderComponent: HeaderComponent,
    HomePostComponent
  },
  async created() {
    this.userid = parseInt(localStorage.getItem('userid'));
    this.username = localStorage.getItem('username');
    this.feeling = parseInt(localStorage.getItem('feeling'));
    this.picture = localStorage.getItem('picture');
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
        :id="1"
        :picture="'http://t0.gstatic.com/licensed-image?q=tbn:ANd9GcRWKkXebtXjMBK3LeRFVhysBnAJB2tVGufs6qZU1Ugw57VoHFu0jPSoukoppxArWzDIJXVyzBNVPDRsU809F_U'"
        :name="'Test123'"
        :feeling="1"
        :profilePicture="'https://wallpapers.com/images/hd/cool-profile-picture-87h46gcobjl5e4xu.jpg'"
        :date="new Date()"
        :caption="'Hello world, this is a test.'"
        :likeCount="123"
      />
      <HomePostComponent
        :id="2"
        :picture="'https://res.cloudinary.com/hello-tickets/image/upload/ar_1:1,c_fill,f_auto,q_auto,w_800/v1643250070/ac5eymqcqcggmafxzljc.jpg'"
        :name="'Test123'"
        :feeling="1"
        :profilePicture="'https://wallpapers.com/images/hd/cool-profile-picture-87h46gcobjl5e4xu.jpg'"
        :date="new Date()"
        :caption="'Hello world, this is a test.'"
        :likeCount="125"
      />
      <HomePostComponent
        :id="3"
        :picture="'http://t0.gstatic.com/licensed-image?q=tbn:ANd9GcRWKkXebtXjMBK3LeRFVhysBnAJB2tVGufs6qZU1Ugw57VoHFu0jPSoukoppxArWzDIJXVyzBNVPDRsU809F_U'"
        :name="'Test123'"
        :feeling="1"
        :profilePicture="'https://wallpapers.com/images/hd/cool-profile-picture-87h46gcobjl5e4xu.jpg'"
        :date="new Date()"
        :caption="'Hello world, this is a test.'"
        :likeCount="10000"
      />

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
