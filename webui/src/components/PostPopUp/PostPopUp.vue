<script>
import PopUpLikeCard from '../PopUpLikeCard.vue';
import PostPopUpCommentCard from './PostPopUpCommentCard.vue';
import HeartIcon from 'vue-material-design-icons/Heart.vue';
import BrokenHeartIcon from 'vue-material-design-icons/HeartBroken.vue';
import SendIcon from 'vue-material-design-icons/Send.vue';
import TrashCanIcon from 'vue-material-design-icons/TrashCan.vue';
import { getPictureURL } from '../../functions/getPictureURL';

export default {
  data() {
    return {
      selectedView: 'comments',
      showHeart: null,
      currentUserID: null,
      token: null,
      likes: null,
      comments: null
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
    updatePost: {
      type: Function,
      required: false
    },
    removePost: {
      type: Function,
      required: false
    }
  },
  methods: {
    getPictureURL,
    async toggleLike() {
      if (this.likes != null) {
        let index = this.likes.findIndex(
          (like) => like.userid == this.currentUserID
        );
        let alreadyLiked = index != -1;

        if (alreadyLiked) {
          await this.$axios
            .delete(`/likes/${this.postid}`, {
              headers: {
                Authorization: `Bearer ${this.token}`
              }
            })
            .then((response) => {
              this.likes.splice(index, 1);
              this.updatePost(this.postid, 'likecount', this.likes.length);
            })
            .catch((error) => {
              alert('Could not unlike post. Please try again.');
            });
        } else {
          await this.$axios
            .put(
              `/likes/${this.postid}`,
              {},
              {
                headers: {
                  Authorization: `Bearer ${this.token}`
                }
              }
            )
            .then((response) => {
              this.likes = [
                {
                  userid: this.currentUserID,
                  username: localStorage.getItem('username'),
                  feeling: localStorage.getItem('feeling'),
                  bio: localStorage.getItem('bio'),
                  picture: localStorage.getItem('picture')
                },
                ...this.likes
              ];
              this.updatePost(this.postid, 'likecount', this.likes.length);
            })
            .catch((error) => {
              alert('Could not like post. Please try again.');
            });
        }

        this.showHeart = alreadyLiked ? 'unlike' : 'like'; // temporarily hard coded
        setTimeout(() => {
          this.showHeart = null;
        }, 1000);
      }
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
            this.comments = [...(this.comments ?? []), comment];
            // clear comment input
            document.getElementById('comment-input').value = '';
            // scroll to bottom of comments
            setTimeout(() => {
              document.getElementById('end-of-comments').scrollIntoView({
                behavior: 'smooth',
                block: 'end',
                inline: 'nearest'
              });
            }, 100);
          })
          .catch((err) => {
            alert("Couldn't send comment. Please try again");
          });
      }
    },
    async downloadLikes() {
      await this.$axios
        .get(`/likes/${this.postid}`, {
          headers: {
            Authorization: `Bearer ${this.token}`
          }
        })
        .then((response) => {
          this.likes = response.data.users ?? [];
        })
        .catch((error) => {
          alert('Could not download likes. Please try again.');
        });
    },
    async downloadComments() {
      await this.$axios
        .get(`/comments/${this.postid}`, {
          headers: {
            Authorization: `Bearer ${this.token}`
          }
        })
        .then((response) => {
          this.comments = response.data.comments ?? [];
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
    },
    async deletePost() {
      if (confirm('Are you sure you want to delete this post?')) {
        this.$axios
          .delete(`/posts/${this.postid}`, {
            headers: {
              Authorization: `Bearer ${this.token}`
            }
          })
          .then((response) => {
            this.removePost(this.postid);
          })
          .catch((error) => {
            alert('Could not delete post. Please try again later.');
          });
      }
    }
  },
  mounted() {
    this.downloadComments();
    this.downloadLikes();
  },
  beforeUnmount() {
    document.body.classList.remove('no-scroll');
  },
  components: {
    PopUpLikeCard,
    PostPopUpCommentCard,
    HeartIcon,
    BrokenHeartIcon,
    SendIcon,
    TrashCanIcon
  },
  created() {
    this.currentUserID = parseInt(localStorage.getItem('userid'));
    this.token = localStorage.getItem('token');
  }
};
</script>

<template>
  <div @click="this.closePost()" class="post-popup-outer-container">
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
              {{ (this.comments?.length ?? 'No') + ' ' }}Comments
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
            <button
              v-if="this.userid == this.currentUserID"
              @click="this.deletePost()"
              class="post-popup-view-option-button"
            >
              <TrashCanIcon />
            </button>
          </div>
          <div
            v-if="this.selectedView == 'likes'"
            class="post-popup-view-section-interactions"
          >
            <PopUpLikeCard
              v-for="like in this.likes"
              :userid="like.userid"
              :key="like.userid"
              :name="like.username"
              :feeling="like.feeling"
              :bio="like.bio"
              :picture="like.picture"
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
                :feeling="this.feeling"
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
              <div id="end-of-comments" />
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
  </div>
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
          overflow-y: scroll;
          height: 100%;
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
