<script>
import PostPopUpLikeCard from './PostPopUpLikeCard.vue';
import PostPopUpCommentCard from './PostPopUpCommentCard.vue';
import HeartIcon from 'vue-material-design-icons/Heart.vue';
import BrokenHeartIcon from 'vue-material-design-icons/HeartBroken.vue';
import SendIcon from 'vue-material-design-icons/Send.vue';

export default {
  data() {
    return {
      selectedView: 'comments',
      showHeart: null
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
    date: {
      type: Date,
      required: true
    },
    caption: {
      type: String,
      required: true
    },
    closePost: {
      type: Function,
      required: true
    },
    togglePostLike: {
      type: Function,
      required: true
    }
  },
  methods: {
    toggleLike() {
      // todo choose correct heart to show (broken or not)
      this.showHeart = 'like'; // temporarily hard coded
      setTimeout(() => {
        this.showHeart = null;
      }, 1000);
      this.togglePostLike();
    },
    sendComment() {
      console.log('send comment');
    }
  },
  mounted() {
    console.log('TODO download comments for post with id: ' + this.id);
  },
  beforeDestroy() {
    document.body.classList.remove('no-scroll');
  },
  components: {
    PostPopUpLikeCard,
    PostPopUpCommentCard,
    HeartIcon,
    BrokenHeartIcon,
    SendIcon
  }
};
</script>

<template>
  <button @click="this.closePost()" class="post-popup-outer-container">
    <div @click.stop="() => {}" class="post-popup-container">
      <div class="post-popup-content">
        <div class="post-popup-image-container">
          <HeartIcon
            :class="{
              'post-popup-image-container-heart': true,
              show: this.showHeart == 'like'
            }"
          />
          <BrokenHeartIcon
            :class="{
              'post-popup-image-container-heart': true,
              show: this.showHeart == 'unlike'
            }"
          />
          <img v-on:dblclick="toggleLike()" :src="this.pictureURL" alt="" />
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
              <!-- todo change this to comment count -->
              {{ '14 ' }}Comments
            </button>
            <button
              @click="this.selectedView = 'likes'"
              :class="{
                'post-popup-view-option-button': true,
                selected: this.selectedView == 'likes'
              }"
            >
              {{ likeCount + ' ' }} Likes
            </button>
          </div>
          <div
            v-if="this.selectedView == 'likes'"
            class="post-popup-view-section-interactions"
          >
            <PostPopUpLikeCard
              v-for="like in [
                1, 2, 3, 4, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
                1, 1, 1, 1, 1, 1, 1, 1
              ]"
              :userid="1"
              :key="like"
              name="Frank123"
              :feeling="1"
              bio="I am a happy person because I am happy and have a happy life."
              pictureUrl="https://pics.craiyon.com/2023-07-15/dc2ec5a571974417a5551420a4fb0587.webp"
            />
          </div>
          <div
            v-if="this.selectedView == 'comments'"
            class="post-popup-view-section-comments-section"
          >
            <div class="post-popup-view-section-interactions">
              <PostPopUpCommentCard
                :userid="1"
                :authorcomment="true"
                :caption="true"
                name="Lorentz27"
                :date="new Date()"
                :feeling="1"
                comment="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla vitae ex auctor, aliquet nisl sed, consequat mi."
                pictureUrl="https://pics.craiyon.com/2023-07-15/dc2ec5a571974417a5551420a4fb0587.webp"
              />
              <PostPopUpCommentCard
                v-for="like in [
                  1, 2, 3, 4, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
                  1, 1, 1, 1, 1, 1, 1, 1
                ]"
                :userid="2"
                :authorcomment="false"
                :key="like"
                :caption="false"
                name="John Doe"
                :feeling="1"
                :date="new Date()"
                comment="I am an AI programming assistant. I can help you with your coding needs. Just ask me anything related to software development."
                pictureUrl="https://pics.craiyon.com/2023-07-15/dc2ec5a571974417a5551420a4fb0587.webp"
              />
            </div>
            <div class="post-popup-comment-input-section">
              <textarea
                v-model="inputValue"
                type="text"
                name="comment-input"
                id="1"
                autogrow="true"
                placeholder="Your comment..."
              />
              <SendIcon
                class="post-popup-comment-send-button"
                style="cursor: pointer"
                @click="sendComment()"
                color="rgb(86, 86, 86)"
                size="30"
              />
            </div>
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
  position: fixed;
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
        position: relative;
        .post-popup-image-container-heart {
          position: absolute;
          top: 50%;
          left: 47%;
          transform: scale(0);
          z-index: 4;
          color: white;
          transition: all 0.2s ease;
          &.show {
            transform: scale(5);
          }
        }
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
          margin-bottom: 30px;
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
              transform: scale(1.02);
            }
            &.selected {
              text-decoration: underline;
            }
          }
        }
        .post-popup-view-section-comments-section {
          display: flex;
          flex-direction: column;
          justify-content: flex-start;
          align-items: center;
          width: 100%;
          padding: 0px 20px;
          overflow-y: scroll;
          .post-popup-comment-input-section {
            display: flex;
            flex-direction: row;
            align-items: center;
            width: 100%;
            min-height: 80px;
            background: white;
            border-radius: 10px;
            margin-top: 10px;
            box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.05);
            textarea {
              width: 88%;
              height: 100%;
              background: none;
              outline: none;
              border: none;
              padding: 5px 10px;
              font-size: 20px;
              font-weight: 600;
              color: rgb(86, 86, 86);
              letter-spacing: 1px;
              font-size: 18px;
              resize: none;
              &::placeholder {
                color: rgb(86, 86, 86);
                font-weight: 600;
                letter-spacing: 1px;
              }
            }
            .post-popup-comment-send-button {
              display: flex;
              align-items: center;
              justify-content: center;
              width: 12%;
              height: 70%;
              color: orange;
              transition: all 0.3s ease;
              border-left: 1px solid rgb(237, 237, 237);
              &:hover {
                transform: scale(1.1);
              }
            }
          }
        }
        .post-popup-view-section-interactions {
          display: flex;
          flex-direction: column;
          justify-content: flex-start;
          align-items: center;
          width: 100%;

          padding: 0px 20px;
          overflow-y: scroll;
        }
      }
    }
  }
}
</style>
