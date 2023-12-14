<script>
import PopUpLikeCard from '../PopUpLikeCard.vue';
import PostPopUpCommentCard from './PostPopUpCommentCard.vue';
import HeartIcon from 'vue-material-design-icons/Heart.vue';
import BrokenHeartIcon from 'vue-material-design-icons/HeartBroken.vue';
import SendIcon from 'vue-material-design-icons/Send.vue';
import { getPictureURL } from '../../functions/getPictureURL';

export default {
  data() {
    return {
      selectedView: 'comments',
      showHeart: null,
      currentUserID: null,
      token: null,

      comments: []
    };
  },
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
    getPictureURL,
    toggleLike() {
      // todo choose correct heart to show (broken or not)
      this.showHeart = 'like'; // temporarily hard coded
      setTimeout(() => {
        this.showHeart = null;
      }, 1000);
      this.togglePostLike();
    },
    async sendComment() {
      const newcomment = document.getElementById('comment-input').value;
      if (newcomment.length > 0) {
        this.$axios
          .put(
            `/comments/${this.postid}`,
            {
              text: newcomment
            },
            {
              headers: {
                Authorization: `Bearer ${this.token}`
              }
            }
          )
          .then((response) => {
            let comment = response.data;
            if (!this.comments) this.comments = [];
            this.comments.push(comment);
            // clear comment input
            document.getElementById('comment-input').value = '';
          })
          .catch((err) => {
            alert("Couldn't send comment. Please try again");
          });
      }
    },
    async downloadComments() {
      await this.$axios
        .get(`/comments/${this.postid}`, {
          headers: {
            Authorization: `Bearer ${this.token}`
          }
        })
        .then((response) => {
          this.comments = response.data.comments;
        })
        .catch((error) => {
          alert('Could not download comments. Please try again.');
        });
    },
    deleteComment(id) {
      this.$axios
        .delete(`/comments/${id}`, {
          headers: {
            Authorization: `Bearer ${this.token}`
          }
        })
        .then((response) => {
          this.comments = this.comments.filter((comment) => comment.id != id);
        })
        .catch((error) => {
          alert('Could not delete comment. Please try again.');
        });
    }
  },
  mounted() {
    this.downloadComments();
  },
  beforeUnmount() {
    document.body.classList.remove('no-scroll');
  },
  components: {
    PopUpLikeCard,
    PostPopUpCommentCard,
    HeartIcon,
    BrokenHeartIcon,
    SendIcon
  },
  created() {
    this.currentUserID = parseInt(localStorage.getItem('userid'));
    this.token = localStorage.getItem('token');
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
          <img
            v-on:dblclick="toggleLike()"
            :src="this.getPictureURL(this.picture)"
            alt=""
          />
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
            <PopUpLikeCard
              v-for="like in [
                1, 2, 3, 4, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
                1, 1, 1, 1, 1, 1, 1, 1
              ]"
              :userid="1"
              :key="like"
              name="Frank123"
              :feeling="1"
              bio="I am a happy person because I am happy and have a happy life."
              :picture="0"
            />
          </div>
          <div
            v-if="this.selectedView == 'comments'"
            class="post-popup-view-section-comments-section"
          >
            <div class="post-popup-view-section-interactions">
              <PostPopUpCommentCard
                :commentid="-1"
                :userid="this.userid"
                :currentUserID="this.currentUserID"
                :authorcomment="true"
                :caption="true"
                :name="this.name"
                :date="new Date(this.date)"
                :feeling="1"
                :comment="this.caption"
                :picture="this.userPicture"
                :deleteComment="() => {}"
              />
              <PostPopUpCommentCard
                v-if="this.comments?.length > 0"
                v-for="comment in this.comments"
                :commentid="comment.id"
                :userid="comment.userid"
                :currentUserID="this.currentUserID"
                :authorcomment="this.userid == comment.userid"
                :key="comment.id"
                :caption="false"
                :name="comment.username"
                :feeling="comment.feeling"
                :date="new Date(comment.createdat)"
                :comment="comment.text"
                :picture="comment.picture"
                :deleteComment="this.deleteComment"
              />
            </div>
            <div class="post-popup-comment-input-section">
              <textarea
                id="comment-input"
                autogrow="true"
                placeholder="Your comment..."
              />
              <SendIcon
                class="post-popup-comment-send-button"
                style="cursor: pointer"
                @click="sendComment()"
                color="rgb(86, 86, 86)"
                :size="30"
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
    cursor: default;
    .post-popup-content {
      height: 100%;
      width: 100%;
      display: flex;
      flex-direction: row;
      @media screen and (max-width: 1100px) {
        flex-direction: column;
        padding-bottom: 30px;
      }
      .post-popup-image-container {
        width: 60%;
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: white;
        box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.05);
        position: relative;
        @media screen and (max-width: 1100px) {
          width: 100%;
          height: 50%;
        }
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
        @media screen and (max-width: 1100px) {
          width: 100%;
          height: 50%;
          margin-left: 0px;
          margin-top: 20px;
        }
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
          min-height: 100%;
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
