import {html, PolymerElement} from '@polymer/polymer';

import {ViewStyle} from './view-style.js';

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

          <dom-repeat items="[[feed]]">
            <template>
              <section>
                <post-card user="[[item.name]]" text="[[item.text]]"></post-card>
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
      feed: {type: Array, value: []}
    };
  }
}

window.customElements.define('feed-view', FeedView);
