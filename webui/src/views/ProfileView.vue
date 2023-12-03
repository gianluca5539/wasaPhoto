<script>
import HeaderComponent from '../components/HeaderComponent.vue';
import { getPictureURL } from '../functions/getPictureURL';

export default {
  data() {
    return {
      profileuserid: null,
      userid: null,
      username: null,
      bio: null,
      feeling: null,
      picture: null
    };
  },
  methods: {
    getPictureURL
  },
  components: { HeaderComponent },
  async created() {
    this.profileuserid = this.$route.params.id;

    // here you should query the stuff needed

    this.userid = parseInt(localStorage.getItem('userid'));
    this.username = localStorage.getItem('username');
    this.feeling = parseInt(localStorage.getItem('feeling'));
    this.bio = localStorage.getItem('bio');
    this.picture = localStorage.getItem('picture');
  }
};
</script>

<template>
  <div class="profile-page-container">
    <HeaderComponent
      :userid="this.userid"
      :username="this.username"
      :feeling="this.feeling"
      :picture="this.picture"
    />
    <div class="profile-page-content">
      <div class="profile-info-container">
        <img :src="getPictureURL(-1)" alt="" />
        <div class="profile-info-data">
          <div class="profile-info-userame">{{ username }}</div>
          <div class="profile-info-bio">{{ bio || 'No bio, yet.' }}</div>
        </div>
      </div>
      <div class="profile-page-posts">
        <button
          v-if="this.profileuserid == this.userid"
          class="profile-page-newpost-card"
        >
          <div class="profile-page-newpost-card-plus">+</div>
          <div class="profile-page-newpost-card-title">New Post</div>
        </button>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.profile-page-container {
  width: 100vw;
  min-height: 100vh;
  height: fit-content;
  display: flex;
  flex-direction: column;
  background: url('../assets/images/homepage-background.jpg') no-repeat fixed;
  .profile-page-content {
    margin-top: 80px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    padding: 0px 20px;
    .profile-info-container {
      width: 80%;
      padding: 18px;
      display: flex;
      flex-direction: row;
      align-self: flex-start;
      align-items: center;
      justify-content: flex-start;
      background-color: white;
      border-radius: 12px;
      margin-bottom: 50px;
      box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.05);
      @media screen and (max-width: 800px) {
        flex-direction: column;
        align-self: center;
      }
      img {
        width: 150px;
        height: 150px;
        min-height: 150px;
        min-width: 150px;
        border-radius: 10%;
        margin-right: 20px;
      }
      .profile-info-data {
        display: flex;
        height: 100%;
        flex-direction: column;
        align-items: flex-start;
        justify-content: flex-start;
        text-overflow: ellipsis;
        overflow: hidden;
        word-wrap: break-word;
        @media screen and (max-width: 800px) {
          align-items: center;
          text-align: center;
        }
        .profile-info-userame {
          font-size: 24px;
          font-weight: bold;
          max-lines: 1;
        }
        .profile-info-bio {
          font-size: 18px;
          font-weight: 300;
          max-lines: 3;
        }
      }
    }
    .profile-page-posts {
      display: flex;
      flex-direction: row;
      flex-wrap: wrap;
      width: 100%;
      padding: 0px 30px;
      align-items: center;
      @media screen and (max-width: 700px) {
        flex-direction: column;
        align-items: center;
      }
      .profile-page-newpost-card {
        width: 300px;
        height: 400px;
        margin-right: 30px;
        margin-bottom: 30px;
        border-radius: 12px;
        border: none;
        box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.05);
        background-color: white;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-around;
        transition: all 0.2s ease-in-out;
        &:hover {
          transform: scale(1.01);
        }
        .profile-page-newpost-card-plus {
          border-radius: 50%;
          background-color: rgb(248, 248, 248);
          color: gray;
          width: 220px;
          height: 220px;
          font-size: 100px;
          display: flex;
          align-items: center;
          justify-content: center;
          text-align: center;
          padding-bottom: 20px;
        }
        .profile-page-newpost-card-title {
          color: gray;
          font-size: 26px;
          font-family: 'Courier New', Courier, monospace;
          font-weight: bold;
        }
      }
    }
  }
}
</style>
