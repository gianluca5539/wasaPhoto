<script>
import HeartIcon from 'vue-material-design-icons/Heart.vue';
import PostPopUp from '../PostPopUp/PostPopUp.vue';
import { getPictureURL } from '../../functions/getPictureURL';

export default {
  name: 'HomePostComponent',
  props: {
    postid: {
      type: Number,
      required: true
    },
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
    picture: {
      type: Number,
      required: true
    },
    userPicture: {
      type: Number,
      required: true
    },
    likeCount: {
      type: Number,
      required: true
    },
    date: {
      type: Number,
      required: true
    },
    caption: {
      type: String,
      required: true
    },
    updatePost: {
      type: Function,
      required: false
    },
    openPost: {
      type: Function,
      required: true
    },
    closePost: {
      type: Function,
      required: true
    },
    postOpen: {
      type: Boolean,
      required: true
    },
    removePost: {
      type: Function,
      required: true
    }
  },
  methods: {
    getPictureURL
  },
  components: {
    HeartIcon,
    PostPopUp
  },
  beforeUnmount() {
    document.body.classList.remove('no-scroll'); // enable scrolling on page
  }
};
</script>

<template>
  <PostPopUp
    v-if="this.postOpen"
    :postid="this.postid"
    :userid="this.userid"
    :name="this.name"
    :feeling="this.feeling"
    :picture="this.picture"
    :userPicture="this.userPicture"
    :likeCount="this.likeCount"
    :date="this.date"
    :caption="this.caption"
    :closePost="this.closePost"
    :updatePost="this.updatePost"
    :removePost="this.removePost"
  />
  <button @click="this.openPost()" class="home-post-card">
    <div class="home-post-card-image-container">
      <img
        :src="this.getPictureURL(this.picture)"
        class="home-post-card-image"
      />
      <div class="home-post-card-likes">
        <HeartIcon class="home-post-card-like-icon" />
        <div class="home-post-card-like-text">{{ likeCount ?? '0' }}</div>
      </div>
    </div>
    <div class="home-post-card-details">
      <div class="home-post-card-user-picture">
        <img :src="this.getPictureURL(this.userPicture)" alt="" />
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
    transform: scale(1.01);
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
  .home-post-card-image-container {
    position: relative;
    width: fit-content;
    height: fit-content;
    margin-bottom: 10px;
    .home-post-card-image {
      width: 100%;
      max-height: 1000px;
      object-fit: cover;
      border-radius: 20px;
      @media screen and (max-width: 650px) {
        border-radius: 0;
      }
    }
    .home-post-card-likes {
      position: absolute;
      bottom: 10px;
      right: 10px;
      background: rgba(0, 0, 0, 0.5);
      color: white;
      padding: 5px 10px;
      border-radius: 10px;
      font-size: 14px;
      font-weight: bold;
      width: fit-content;
      display: flex;
      flex-direction: row;
      border: none;
      .home-post-card-like-icon {
        margin-right: 5px;
      }
      .home-post-card-like-text {
        color: white !important;
      }
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
