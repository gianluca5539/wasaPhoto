<script>
export default {
  name: 'HomePostComponent',
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
    },
    profilePictureURL: {
      type: String,
      required: true
    },
    caption: {
      type: String,
      required: true
    }
  }
}
</script>

<template>
  <button class="home-post-card">
    <img :src="pictureURL" class="home-post-card-image" />
    <div class="home-post-card-details">
      <div class="home-post-card-user-picture">
        <img :src="profilePictureURL" alt="" />
        <div class="home-post-card-user-picture-feeling">
          <span v-if="feeling === 0">😐</span>
          <span v-if="feeling === 1">😀</span>
          <span v-if="feeling === 2">😍</span>
          <span v-if="feeling === 3">😡</span>
          <span v-if="feeling === 4">😭</span>
        </div>
      </div>
      <div class="home-post-card-caption-container">
        <div class="home-post-card-caption-username">{{ name }}</div>
        <div class="home-post-card-caption">
          {{ caption }}
        </div>
      </div>
    </div>
  </button>
</template>

<style lang="scss">
.home-post-card {
  margin-top: 20px;
  margin-bottom: 10px;
  margin-left: 10px;
  margin-right: 10px;
  width: 600px;
  min-width: 600px;
  height: fit-content;
  display: flex;
  flex-direction: column;
  padding: 10px;
  border-radius: 20px;
  border: none;
  box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.15);
  background: white;
  position: relative;
  transition: all 0.3s ease;
  &:hover {
    transform: scale(1.02);
    cursor: pointer;
  }
  @media screen and (max-width: 650px) {
    width: 100%;
    border-radius: 0;
    margin-left: 0;
    margin-right: 0;
    box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.15);
    min-width: 0;
  }
  .home-post-card-image {
    width: 100%;
    max-height: 1000px;
    object-fit: cover;
    border-radius: 20px;
    margin-bottom: 10px;
    @media screen and (max-width: 650px) {
      border-radius: 0;
    }
  }
  .home-post-card-details {
    display: flex;
    flex-direction: row;

    padding: 0px 10px;
    .home-post-card-user-picture {
      display: flex;
      flex-direction: row;
      align-items: center;
      margin-right: 20px;
      position: relative;
      width: 50px;
      height: 50px;
      min-width: 50px;
      min-height: 50px;
      img {
        object-fit: cover;
        overflow: hidden;
        width: 100%;
        height: 100%;
        border-radius: 50%;
      }
      .home-post-card-user-picture-feeling {
        position: absolute;
        bottom: -5px;
        right: -5px;
        font-size: 20px;
      }
    }
    .home-post-card-caption-container {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      width: 85%;
      .home-post-card-caption-username {
        font-size: 18px;
        font-weight: bold;
        font-family: 'Courier New', Courier, monospace;
      }
      .home-post-card-caption {
        font-size: 16px;
        font-family: 'Courier New', Courier, monospace;
        text-align: justify;

        word-wrap: break-word;
        @media screen and (max-width: 650px) {
          font-size: 16px;
        }
      }
    }
  }
}
</style>
