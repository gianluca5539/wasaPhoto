<script>
export default {
  props: {
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
    pictureUrl: {
      type: String,
      required: true
    }
  }
};
</script>

<template>
  <button class="post-popup-like-card">
    <div class="post-popup-like-card-picure">
      <img :src="pictureUrl" alt="Profile Picture" />
      <div class="post-popup-like-card-feeling">
        <span v-if="feeling === 0">😐</span>
        <span v-if="feeling === 1">😀</span>
        <span v-if="feeling === 2">😍</span>
        <span v-if="feeling === 3">😡</span>
        <span v-if="feeling === 4">😭</span>
      </div>
    </div>
    <div class="post-popup-like-card-details">
      <div class="post-popup-like-card-name">{{ name }}</div>
      <div class="post-popup-like-card-bio">{{ bio }}</div>
    </div>
  </button>
</template>

<style lang="scss">
.post-popup-like-card {
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
  .post-popup-like-card-picure {
    width: 50px;
    height: 50px;
    position: relative;
    img {
      width: 60px;
      height: 60px;
      border-radius: 50%;
    }
    .post-popup-like-card-feeling {
      font-size: 22px;
      position: absolute;
      bottom: -14px;
      right: -12px;
    }
  }
  .post-popup-like-card-details {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
    margin-left: 25px;
    max-width: 100%;
    overflow: hidden;
    .post-popup-like-card-name {
      font-size: 18px;
      font-weight: 600;
      color: rgb(50, 50, 50);
      max-width: 100%;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
    .post-popup-like-card-bio {
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
