import { html } from '@polymer/lit-element';

import { PageViewElement } from './page-view-element.js';
import { SharedStyles } from './shared-styles.js';

import './post-card.js';

/**
 * @customElement
 * @polymer
 */
class FeedView extends PageViewElement {
  render() {
    return html`
      ${SharedStyles}
      <style>
        ul {
          list-style: none;
        }
      </style>

      <section>
        <h2>Feed</h2>
        <ul>
          ${this.posts.map((i) => html`<li><post-card .user="${i}"></post-card></li>`)}
        </ul>
      </section>
    `;
  }
  static get properties() {
    return {
      posts: {type: Array}
    };
  }
  constructor() {
    super();
    this.posts = [1,2,3,4,5,6,7,8,9,0];
  }
}

window.customElements.define('feed-view', FeedView);
