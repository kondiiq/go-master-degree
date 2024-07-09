<template>
    <div v-if="posts">
      <v-data-table
      v-model:items-per-page="itemsPerPage"
      :headers="headers"
      :items="data" item-value="userId"
      class="elevation-1"></v-data-table>
    </div>
  </template>
  <script>
  import { VDataTable } from 'vuetify/labs/VDataTable'
  import getAllPosts from '../services/getAllPosts'
  
  export default {
    components: {
      VDataTable,
    },
    setup() {
  
      const { oPosts, oError, fnLoad } = getAllPosts()
      fnLoad();
      console.log(oPosts.value);
      return {
        oPosts,
        oError,
        itemsPerPage: 10,
        headers: [
          {
            title: 'User Id',
            align: 'start',
            sortable: false,
            key: 'userId',
          },
          { title: 'Title', align: 'start', key: 'title' },
          { title: 'Body', align: 'start', key: 'body' },
        ],
        data: oPosts,
      }
  
    },
  }
  </script>
  <style lang="">
  
  </style>