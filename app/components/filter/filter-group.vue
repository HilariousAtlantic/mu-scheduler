<template>

  <div class="filter-group">

    <button class="group-header btn btn-block" @click="$emit('edit')">

      <span>{{label}}</span>

      <span v-if="editing" class="fa fa-caret-up"></span>
      <span v-else class="fa fa-caret-down"></span>

    </button>

    <div v-if="editing" class="group-editor" :style="editorPosition">

      <slot></slot>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'filter-group',

    props: ['label', 'index', 'editing'],

    computed: {

      editorPosition() {

        let percent = this.index*100;

        if (percent < 50) {

          return {left: -percent + '%'};

        } else if (percent > 50) {

          return {right: percent-100 + '%'};

        } else {

          return {left: '50%', right: '50%', transform: 'translateX(-50%)'};

        }

      }

    }

  }

</script>

<style scoped>

  .filter-group {
    position: relative;
  }

  .group-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .group-editor {
    position: absolute;
    width: 450px;
    padding: 20px;
    background: #fff;
    border: 1px solid #ddd;
    border-radius: 3px;
    box-shadow: 0 3px 12px rgba(27,31,35,0.15);
    top: calc(100% + 5px);
    z-index: 10;
  }

</style>
