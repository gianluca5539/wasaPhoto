<script>
import HomeHeaderComponent from '../components/HomeView/HomeHeaderComponent.vue'
import HomePostComponent from '../components/HomeView/HomePostComponent.vue'
export default {
  components: {
    HomeHeaderComponent: HomeHeaderComponent,
    HomePostComponent
  }
}
</script>

<template>
  <div class="homepage-container">
    <HomeHeaderComponent
      :name="'Test123'"
      :feeling="1"
      :pictureURL="'https://wallpapers.com/images/hd/cool-profile-picture-87h46gcobjl5e4xu.jpg'"
    />
    <div class="homepage-feed-container">
      <HomePostComponent
        :pictureURL="'http://t0.gstatic.com/licensed-image?q=tbn:ANd9GcRWKkXebtXjMBK3LeRFVhysBnAJB2tVGufs6qZU1Ugw57VoHFu0jPSoukoppxArWzDIJXVyzBNVPDRsU809F_U'"
        :name="'Test123'"
        :feeling="1"
        :profilePictureURL="'https://wallpapers.com/images/hd/cool-profile-picture-87h46gcobjl5e4xu.jpg'"
        :caption="'Hello world, this is a test.'"
      />
      <HomePostComponent
        :pictureURL="'https://res.cloudinary.com/hello-tickets/image/upload/ar_1:1,c_fill,f_auto,q_auto,w_800/v1643250070/ac5eymqcqcggmafxzljc.jpg'"
        :name="'Test123'"
        :feeling="1"
        :profilePictureURL="'https://wallpapers.com/images/hd/cool-profile-picture-87h46gcobjl5e4xu.jpg'"
        :caption="'Hello world, this is a test.'"
      />
      <HomePostComponent
        :pictureURL="'http://t0.gstatic.com/licensed-image?q=tbn:ANd9GcRWKkXebtXjMBK3LeRFVhysBnAJB2tVGufs6qZU1Ugw57VoHFu0jPSoukoppxArWzDIJXVyzBNVPDRsU809F_U'"
        :name="'Test123'"
        :feeling="1"
        :profilePictureURL="'https://wallpapers.com/images/hd/cool-profile-picture-87h46gcobjl5e4xu.jpg'"
        :caption="'Hello world, this is a test.'"
      />
      <HomePostComponent
        :pictureURL="'http://t0.gstatic.com/licensed-image?q=tbn:ANd9GcRWKkXebtXjMBK3LeRFVhysBnAJB2tVGufs6qZU1Ugw57VoHFu0jPSoukoppxArWzDIJXVyzBNVPDRsU809F_U'"
        :name="'Test123'"
        :feeling="1"
        :profilePictureURL="'https://wallpapers.com/images/hd/cool-profile-picture-87h46gcobjl5e4xu.jpg'"
        :caption="'Hello world, this is a test.'"
      />
      <HomePostComponent
        :pictureURL="'http://t0.gstatic.com/licensed-image?q=tbn:ANd9GcRWKkXebtXjMBK3LeRFVhysBnAJB2tVGufs6qZU1Ugw57VoHFu0jPSoukoppxArWzDIJXVyzBNVPDRsU809F_U'"
        :name="'Test123'"
        :feeling="1"
        :profilePictureURL="'https://wallpapers.com/images/hd/cool-profile-picture-87h46gcobjl5e4xu.jpg'"
        :caption="'Hello world, this is a test.'"
      />
    </div>
  </div>
</template>

<style lang="scss">
.homepage-container {
  width: 100vw;
  max-width: 100vw;
  display: flex;
  flex-direction: column;
  background: url('../assets/images/homepage-background.jpg') no-repeat center center fixed;
  .homepage-feed-container {
    margin-top: 20px;
    display: flex;
    flex-direction: column;
    width: 100%;
    height: fit-content;
    align-items: center;
  }
}
</style>
