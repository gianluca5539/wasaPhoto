<script>
import UserHeaderCard from './UserHeaderCard.vue';

export default {
  name: 'HeaderComponent',
  props: {
    userid: {
      type: Number,
      required: true
    },
    username: {
      type: String,
      required: true
    },
    feeling: {
      type: Number,
      required: true
    },
    picture: {
      type: Number,
      required: true
    }
  },
  methods: {
    goToHomePage() {
      this.$router.push('/');
    }
  },
  components: { UserHeaderCard }
};
</script>

<template>
  <div class="homepage-header-container">
    <button @click="goToHomePage()" class="homepage-header-title">
      WASAPhoto
    </button>
    <UserHeaderCard
      :userid="userid"
      :username="username"
      :feeling="feeling"
      :picture="picture"
    />
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
    background: transparent;
    border: none;
    outline: none;
    cursor: pointer;
    transition: all 0.3s ease;
    &:hover {
      color: orange;
      transform: scale(1.01);
    }
  }
}
</style>
