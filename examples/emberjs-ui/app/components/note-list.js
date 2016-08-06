import Ember from 'ember';

export default Ember.Component.extend({
  init(){
    this._super(...arguments);
    this.set('notes', []);
  },

  actions: {
    saveNote(){
      var note = this.$('#addNote :input[id="note"]').val();
      this.get('notes').push(note);
      this.notifyPropertyChange('notes');
    }
  }
});
