/**
 * AYLIEN News API
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * The version of the OpenAPI document: 3.0
 * Contact: support@aylien.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The RepresentativeStory model module.
 * @module model/RepresentativeStory
 * @version 4.5.0
 */
class RepresentativeStory {
    /**
     * Constructs a new <code>RepresentativeStory</code>.
     * @alias module:model/RepresentativeStory
     */
    constructor() { 
        
        RepresentativeStory.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>RepresentativeStory</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RepresentativeStory} obj Optional instance to populate.
     * @return {module:model/RepresentativeStory} The populated <code>RepresentativeStory</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RepresentativeStory();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');

                if ('id' !== 'id') {
                  Object.defineProperty(obj, 'id', {
                    get() {
                      return obj['id'];
                    }
                  });
                }
            }
            if (data.hasOwnProperty('permalink')) {
                obj['permalink'] = ApiClient.convertToType(data['permalink'], 'String');

                if ('permalink' !== 'permalink') {
                  Object.defineProperty(obj, 'permalink', {
                    get() {
                      return obj['permalink'];
                    }
                  });
                }
            }
            if (data.hasOwnProperty('published_at')) {
                obj['published_at'] = ApiClient.convertToType(data['published_at'], 'Date');

                if ('published_at' !== 'publishedAt') {
                  Object.defineProperty(obj, 'publishedAt', {
                    get() {
                      return obj['published_at'];
                    }
                  });
                }
            }
            if (data.hasOwnProperty('title')) {
                obj['title'] = ApiClient.convertToType(data['title'], 'String');

                if ('title' !== 'title') {
                  Object.defineProperty(obj, 'title', {
                    get() {
                      return obj['title'];
                    }
                  });
                }
            }
        }
        return obj;
    }


}

/**
 * ID of the story which is a unique identification
 * @member {Number} id
 */
RepresentativeStory.prototype['id'] = undefined;

/**
 * The story permalink URL
 * @member {String} permalink
 */
RepresentativeStory.prototype['permalink'] = undefined;

/**
 * Published date of the story
 * @member {Date} published_at
 */
RepresentativeStory.prototype['published_at'] = undefined;

/**
 * Title of the story
 * @member {String} title
 */
RepresentativeStory.prototype['title'] = undefined;






export default RepresentativeStory;

