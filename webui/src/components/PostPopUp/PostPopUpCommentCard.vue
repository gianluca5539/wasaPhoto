<script>
import { getPictureURL } from '../../functions/getPictureURL';
import TrashCanIcon from 'vue-material-design-icons/TrashCan.vue';
export default {
  props: {
    commentid: {
      type: Number,
      required: true
    },
    userid: {
      type: Number,
      required: true
    },
    currentUserID: {
      type: Number,
      required: true
    },
    name: {
      type: String,
      required: true
    },
    authorcomment: {
      type: Boolean,
      required: true
    },
    caption: {
      type: Boolean,
      required: true
    },
    date: {
      type: Date,
      required: true
    },
    feeling: {
      type: Number,
      required: true
    },
    comment: {
      type: String,
      required: true
    },
    picture: {
      type: Number,
      required: true
    },
    deleteComment: {
      type: Function,
      required: true
    }
  },
  components: {
    TrashCanIcon
  },
  methods: {
    getPictureURL,
    openUserProfile() {
      this.$router.push('/profile/' + this.userid);
    },
    getDateString() {
      return `${this.date.getDate()}/${this.date.getMonth()}/${this.date.getFullYear()}
          ${this.date.getHours()}:${
        this.date.getMinutes() < 10 ? '0' : ''
      }${this.date.getMinutes()}`;
    }
  }
};
</script>

<template>
  <button
    @click="this.openUserProfile()"
    :class="{
      'post-popup-comment-card': true,
      authorcomment: authorcomment,
      caption: caption
    }"
  >
    <div class="post-popup-comment-card-picure">
      <img :src="this.getPictureURL(this.picture)" alt="Profile Picture" />
      <div class="post-popup-comment-card-feeling">
        <span v-if="this.feeling === 0">😐</span>
        <span v-if="this.feeling === 1">😀</span>
        <span v-if="this.feeling === 2">😍</span>
        <span v-if="this.feeling === 3">😡</span>
        <span v-if="this.feeling === 4">😭</span>
      </div>
    </div>
    <div
      :class="{
        'post-popup-comment-card-details': true,
        authorcomment: authorcomment
      }"
    >
      <div
        :class="{
          'post-popup-comment-card-name': true,
          authorcomment: authorcomment
        }"
      >
        <div>
          {{ this.userid == this.currentUserID ? 'You' : this.name }}
        </div>
        <div class="post-popup-comment-card-date">
          {{ this.getDateString() }}
        </div>
      </div>
      <div
        :class="{
          'post-popup-comment-card-comment': true,
          authorcomment: authorcomment
        }"
      >
        {{ this.comment
        }}
        <button
          v-if="this.userid == this.currentUserID && !this.caption"
          @click.stop="this.deleteComment(this.commentid)"
          class="post-popup-comment-card-delete-button"
        >
          <TrashCanIcon />
        </button>
      </div>
    </div>
  </button>
</template>

<style lang="scss">
.post-popup-comment-card {
  display: flex;
  flex-direction: row-reverse;
  width: 100%;
  min-height: fit-content;
  margin-top: 20px;
  background-color: transparent;
  border: none;
  padding-bottom: 15px;
  border-bottom: 1px solid rgb(241, 241, 241);
  transition: all 0.2s ease-in-out;
  &.authorcomment {
    flex-direction: row;
  }
  &.caption {
    background-color: white;
    padding: 20px;
    box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.05);
  }
  &:hover {
    transform: scale(1.02);
  }
  .post-popup-comment-card-picure {
    width: 50px;
    height: 50px;
    position: relative;
    img {
      width: 60px;
      height: 60px;
      border-radius: 50%;
    }
    .post-popup-comment-card-feeling {
      font-size: 22px;
      position: absolute;
      bottom: -14px;
      right: -12px;
    }
  }
  .post-popup-comment-card-details {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-end;
    margin-right: 25px;
    max-width: 100%;
    min-height: fit-content;
    &.authorcomment {
      align-items: flex-start;
      margin-right: 0px;
      margin-left: 25px;
    }
    .post-popup-comment-card-name {
      font-size: 18px;
      font-weight: 600;
      color: rgb(50, 50, 50);
      max-width: 100%;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      min-height: 20px;
      text-align: right;
      &.authorcomment {
        text-align: left;
        color: orange;
      }

      .post-popup-comment-card-date {
        font-size: 12px;
        font-weight: 100;
        color: rgb(145, 145, 145);
        max-width: 100%;
        word-wrap: break-word;
        text-align: right;
        min-height: fit-content;
        &.authorcomment {
          text-align: left;
        }
      }
    }
    .post-popup-comment-card-comment {
      font-size: 16px;
      font-weight: 600;
      color: rgb(100, 100, 100);
      max-width: 100%;
      word-wrap: break-word;
      text-align: right;
      min-height: fit-content;
      &.authorcomment {
        text-align: left;
      }
      .post-popup-comment-card-delete-button {
        color: rgb(208, 208, 208);
        border: none;
        width: 30px;
        outline: none;
        background-color: transparent;
        transition: all 0.2s ease-in-out;
        &:hover {
          color: red;
        }
      }
    }
  }
}
</style>
