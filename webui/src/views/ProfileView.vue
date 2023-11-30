<script>
import HeaderComponent from '../components/HeaderComponent.vue';

export default {
  data() {
    return {
      profileuserid: null,
      userid: null,
      username: null,
      feeling: null,
      picture: null
    };
  },
  components: { HeaderComponent },
  async created() {
    this.profileuserid = this.$route.params.id;

    // here you should query the stuff needed

    this.userid = parseInt(localStorage.getItem('userid'));
    this.username = localStorage.getItem('username');
    this.feeling = parseInt(localStorage.getItem('feeling'));
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
      <h1>{{ this.profileuserid }}'s Profile</h1>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.profile-page-container {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  background: url('../assets/images/homepage-background.jpg') no-repeat;
  .profile-page-content {
    margin-top: 80px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    padding: 0px 20px;
  }
}
</style>
