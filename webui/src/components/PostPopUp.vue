<script>
export default {
  data() {
    return {
      selectedView: 'comments'
    };
  },
  props: {
    id: {
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
    pictureURL: {
      type: String,
      required: true
    },
    profilePictureURL: {
      type: String,
      required: true
    },
    likeCount: {
      type: Number,
      required: true
    },
    caption: {
      type: String,
      required: true
    },
    closePost: {
      type: Function,
      required: true
    }
  },
  methods: {
    togglePostLike() {
      console.log('like post');
    }
  },
  mounted() {
    console.log('TODO download comments for post with id: ' + this.id);
  }
};
</script>

<template>
  <button @click="this.closePost()" class="post-popup-outer-container">
    <div @click.stop="() => {}" class="post-popup-container">
      <div class="post-popup-content">
        <div class="post-popup-image-container">
          <img :src="pictureURL" alt="" />
        </div>
        <div class="post-popup-info-section">
          <div class="post-popup-view-options">
            <button
              @click="this.selectedView = 'comments'"
              :class="{
                'post-popup-view-option-button': true,
                selected: this.selectedView == 'comments'
              }"
            >
              Comments
            </button>
            <button
              @click="this.selectedView = 'likes'"
              :class="{
                'post-popup-view-option-button': true,
                selected: this.selectedView == 'likes'
              }"
            >
              Likes
            </button>
          </div>
        </div>
      </div>
    </div>
  </button>
</template>

<style lang="scss">
.post-popup-outer-container {
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  top: 0;
  left: 0;
  z-index: 2;
  font-family: 'Courier New', Courier, monospace;
  border: none;
  cursor: default;
  .post-popup-container {
    width: 80vw;
    height: 90vh;
    background: whitesmoke;
    border-radius: 10px;
    padding: 20px 20px;
    z-index: 3;
    cursor: default;
    .post-popup-content {
      height: 100%;
      width: 100%;
      display: flex;
      flex-direction: row;
      .post-popup-image-container {
        width: 60%;
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: white;
        box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.05);
        img {
          width: 100%;
          height: 100%;
          object-fit: contain;
          border-radius: 10px;
        }
      }
      .post-popup-info-section {
        width: 40%;
        height: 100%;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;
        margin-left: 20px;
        .post-popup-view-options {
          display: flex;
          justify-content: center;
          align-items: center;
          .post-popup-view-option-button {
            background: none;
            outline: none;
            border: none;
            cursor: pointer;
            margin: none;
            transition: all 0.3s ease;
            margin: 0 20px;
            font-size: 20px;
            font-weight: 600;
            color: rgb(86, 86, 86);
            letter-spacing: 1px;
            &:hover {
              transform: scale(1.05);
            }
            &.selected {
              text-decoration: underline;
            }
          }
        }
      }
    }
  }
}
</style>
