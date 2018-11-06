import { html, PolymerElement } from '@polymer/polymer';

import { ViewStyle } from './view-style.js';

import './post-card.js';

/**
 * @customElement
 * @polymer
 */
class FeedView extends PolymerElement {
  static get template() {
    return html`
      <dom-if if=[[active]]>
        <template>
          ${ViewStyle}

          <dom-repeat items="[[posts]]">
            <template>
              <section>
                <post-card user="[[item]]"></post-card>
              </section>
            </template>
          </dom-repeat>
        </template>
      </dom-if>
    `;
  }

  static get properties() {
    return {
      active: {type: Boolean, value: false},
      posts: {type: Array, value: [1,2,3,4,5,6,7,8,9,0]}
    };
  }
}

window.customElements.define('feed-view', FeedView);
