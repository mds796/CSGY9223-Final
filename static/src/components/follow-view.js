import {html, PolymerElement} from "../../node_modules/@polymer/polymer/polymer-element.js";
import {ViewStyle} from './view-style.js';
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
              cursor: pointer;
            }
          </style>

          <section>
            <label for="query">Type a search query, then press Enter to search for users:</label>
            <input id="query" type="text" value="{{query::change}}" placeholder="Username"/>
          </section>
            
            <dom-repeat items="[[follows]]" itemsIndexAs="index">
              <template>
                <section>                  
                    <label for="[[item.Name]]">
                        <input type="checkbox" id="[[item.Name]]" name="[[item.Name]]" value="[[index]]" checked$="[[item.Follow]]" on-change="toggleFollow">
                        <user-card user="[[item.name]]"></user-card>
                    </label>
                </section>
              </template>
            </dom-repeat>
        </template>
      </dom-if>
    `;
    }

    static get properties() {
        return {
            active: {
                type: Boolean,
                value: false
            },
            query: {
                type: String,
                value: "",
                observer: "fetchFollows"
            },
            follows: {
                type: Array,
                value: []
            }
        };
    }

    fetchFollows(query) {
        const provider = this;

        fetch('/follows?query=' + query).then(response => {
            if (response.ok) {
                return response.json();
            } else {
                return {};
            }
        }).then(data => {
            provider.follows = data.follows;
        }).catch(err => {
            console.log("Unable to fetch follows: ", err);
        });
    }

    toggleFollow(e) {
        const index = e.target.value;
        const path = 'follows.' + index + '.Follow';

        this.set(path, !this.follows[index].Follow);

        const body = "name=" + e.target.name + "&follow=" + this.follows[index].Follow;

        fetch('/follow', {method: 'POST', body: body, headers:{'Content-Type': 'application/x-www-form-urlencoded'}}).catch(err => {
            console.log("Unable to toggle follow: ", err);
        });
    }
}

window.customElements.define('follow-view', FollowView);