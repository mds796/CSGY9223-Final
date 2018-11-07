import { html, PolymerElement } from '@polymer/polymer';

import { minusIcon } from './icons.js';
import { ViewStyle } from './view-style.js';
import './user-card.js';

/**
 * @customElement
 * @polymer
 */
class FollowView extends PolymerElement {
  static get template() {
    return html`
      <dom-if if=[[active]]>
        <template>
          ${ViewStyle}

          <style>
            user-card {
              display: inline-block;
            }
          </style>

          <section>
            <h2>Follow</h2>
            <input type="text" value="{{query}}" placeholder="Username"/>

            <dom-repeat items="[[users]]">
              <template>
                <section>
                  <button class="remove">${minusIcon}</button>
                  <user-card user="[[item.name]]"></user-card>
                </section>
              </template>
          </dom-repeat>
          </section>
        </template>
      </dom-if>
    `;
  }

  static get properties() {
    return {
      active: {type: Boolean, value: false},
      query: {type: String, value: ""},
      users: {type: Array, computed: '_queryUsers(query)'}
    };
  }

  _queryUsers(query) {
    // call web and update users
    return [{name: 'fake123-' + query}]
  }
}

window.customElements.define('follow-view', FollowView);
