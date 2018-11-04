import { html } from '@polymer/lit-element';

import { PageViewElement } from './page-view-element.js';
import { SharedStyles } from './shared-styles.js';

/**
 * @customElement
 * @polymer
 */
class FeedView extends PageViewElement {
  render() {
    return html`
      ${SharedStyles}
      <h2>Hello ${this.prop1}!</h2>
    `;
  }
  static get properties() {
    return {
      prop1: {
        type: String,
        value: 'feed-view'
      }
    };
  }
  constructor() {
    super();
    this.prop1 = 'Feed View';
  }
}

window.customElements.define('feed-view', FeedView);
