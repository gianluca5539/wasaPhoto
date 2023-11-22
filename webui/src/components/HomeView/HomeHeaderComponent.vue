<script>
import UserHeaderCard from '../UserHeaderCard.vue';

export default {
  name: 'HomeHeaderComponent',
  props: {
    name: {
      type: String,
      required: true
    },
    feeling: {
      type: Number,
      required: true
    },
    pictureURL: {
      type: String,
      required: true
    }
  },
  components: { UserHeaderCard }
};
</script>

<template>
  <div class="homepage-header-container">
    <div class="homepage-header-title">WASAPhoto</div>
    <UserHeaderCard :name="name" :feeling="feeling" :pictureURL="pictureURL" />
  </div>
</template>

<style lang="scss">
.homepage-header-container {
  width: 100%;
  max-width: 100%;
  height: 70px;
  font-family: 'Courier New', Courier, monospace;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 30px;
  position: absolute;
  top: 0px;
  right: 0px;
  .homepage-header-title {
    font-size: 45px;
    font-weight: bold;
    color: white;
  }
}
</style>
