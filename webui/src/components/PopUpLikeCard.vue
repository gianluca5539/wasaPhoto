<script>
import { getPictureURL } from '../functions/getPictureURL';

export default {
  props: {
    userid: {
      type: Number,
      required: true
    },
    name: {
      type: String,
      required: true
    },
    feeling: {
      type: Number,
      required: true
    },
    bio: {
      type: String,
      required: true
    },
    picture: {
      type: Number,
      required: true
    }
  },
  methods: {
    getPictureURL,
    openUserProfile() {
      this.$router.push(`/profile/${this.userid}`);
    }
  }
};
</script>

<template>
  <button @click="this.openUserProfile()" class="popup-like-card">
    <div class="popup-like-card-picure">
      <img :src="this.getPictureURL(this.picture)" alt="Profile Picture" />
      <div class="popup-like-card-feeling">
        <span v-if="feeling === 0">😐</span>
        <span v-if="feeling === 1">😀</span>
        <span v-if="feeling === 2">😍</span>
        <span v-if="feeling === 3">😡</span>
        <span v-if="feeling === 4">😭</span>
      </div>
    </div>
    <div class="popup-like-card-details">
      <div class="popup-like-card-name">{{ name }}</div>
      <div class="popup-like-card-bio">{{ bio || 'No bio, yet.' }}</div>
    </div>
  </button>
</template>

<style lang="scss">
.popup-like-card {
  display: flex;
  flex-direction: row;
  width: 100%;
  height: 80px;
  min-height: 80px;
  margin-top: 20px;
  background-color: transparent;
  border: none;
  border-bottom: 1px solid rgb(231, 231, 231);
  transition: all 0.2s ease-in-out;
  &:hover {
    transform: scale(1.02);
  }
  .popup-like-card-picure {
    width: 50px;
    height: 50px;
    position: relative;
    img {
      width: 60px;
      height: 60px;
      border-radius: 50%;
    }
    .popup-like-card-feeling {
      font-size: 22px;
      position: absolute;
      bottom: -14px;
      right: -12px;
    }
  }
  .popup-like-card-details {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
    margin-left: 25px;
    max-width: 100%;
    overflow: hidden;
    .popup-like-card-name {
      font-size: 18px;
      font-weight: 600;
      color: rgb(50, 50, 50);
      max-width: 100%;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
    .popup-like-card-bio {
      font-size: 14px;
      font-weight: 400;
      color: rgb(100, 100, 100);
      max-width: 100%;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  }
}
</style>
