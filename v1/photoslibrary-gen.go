package photoslibrary // import "google.golang.org/api/photoslibrary/v1"

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
)

const apiId = "photoslibrary:v1"
const apiName = "photoslibrary"
const apiVersion = "v1"
const basePath = "https://photoslibrary.googleapis.com/"
const apiRevision = "20190816"

// OAuth2 scopes used by this API.
const (

	// View and manage your Google Photos library
	PhotoslibraryScope = "https://www.googleapis.com/auth/photoslibrary"

	// Add to your Google Photos library
	PhotoslibraryAppendonlyScope = "https://www.googleapis.com/auth/photoslibrary.appendonly"

	// View your Google Photos library
	PhotoslibraryReadonlyScope = "https://www.googleapis.com/auth/photoslibrary.readonly"

	// Manage photos added by this app
	PhotoslibraryReadonlyAppcreateddataScope = "https://www.googleapis.com/auth/photoslibrary.readonly.appcreateddata"

	// Manage and add to shared albums on your behalf
	PhotoslibrarySharingScope = "https://www.googleapis.com/auth/photoslibrary.sharing"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}

	s.Albums = NewAlbumsService(s)

	s.MediaItems = NewMediaItemsService(s)

	s.SharedAlbums = NewSharedAlbumsService(s)

	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Albums *AlbumsService

	MediaItems *MediaItemsService

	SharedAlbums *SharedAlbumsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAlbumsService(s *Service) *AlbumsService {
	rs := &AlbumsService{s: s}
	return rs
}

type AlbumsService struct {
	s *Service
}

func NewMediaItemsService(s *Service) *MediaItemsService {
	rs := &MediaItemsService{s: s}
	return rs
}

type MediaItemsService struct {
	s *Service
}

func NewSharedAlbumsService(s *Service) *SharedAlbumsService {
	rs := &SharedAlbumsService{s: s}
	return rs
}

type SharedAlbumsService struct {
	s *Service
}

// Request to add an enrichment to a specific album at a specific position.
type AddEnrichmentToAlbumRequest struct {
	// The position in the album where the enrichment is to be inserted.
	AlbumPosition *AlbumPosition `json:"albumPosition,omitempty" `
	// The enrichment to be added.
	NewEnrichmentItem *NewEnrichmentItem `json:"newEnrichmentItem,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *AddEnrichmentToAlbumRequest) MarshalJSON() ([]byte, error) {
	type NoMethod AddEnrichmentToAlbumRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// The enrichment item that's created.
type AddEnrichmentToAlbumResponse struct {
	// Output only. Enrichment which was added.
	EnrichmentItem *EnrichmentItem `json:"enrichmentItem,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *AddEnrichmentToAlbumResponse) MarshalJSON() ([]byte, error) {
	type NoMethod AddEnrichmentToAlbumResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Representation of an album in Google Photos.
// Albums are containers for media items. If an album has been shared by the
// application, it contains an extra `shareInfo` property.
type Album struct {
	// [Output only] A URL to the cover photo's bytes. This shouldn't be used as
	// is. Parameters should be appended to this URL before use. See the
	// [developer
	// documentation](https://developers.google.com/photos/library/guides/access-media-items#base-urls)
	// for a complete list of supported parameters. For example,
	// `'=w2048-h1024'` sets the dimensions of the cover photo to have a width of
	// 2048 px and height of 1024 px.
	CoverPhotoBaseUrl string `json:"coverPhotoBaseUrl,omitempty" `
	// [Output only] Identifier for the media item associated with the cover
	// photo.
	CoverPhotoMediaItemId string `json:"coverPhotoMediaItemId,omitempty" `
	// [Ouput only] Identifier for the album. This is a persistent identifier that
	// can be used between sessions to identify this album.
	Id string `json:"id,omitempty" `
	// [Output only] True if you can create media items in this album.
	// This field is based on the scopes granted and permissions of the album. If
	// the scopes are changed or permissions of the album are changed, this field
	// is updated.
	IsWriteable bool `json:"isWriteable,omitempty" `
	// [Output only] The number of media items in the album.
	MediaItemsCount string `json:"mediaItemsCount,omitempty" `
	// [Output only] Google Photos URL for the album. The user needs to be signed
	// in to their Google Photos account to access this link.
	ProductUrl string `json:"productUrl,omitempty" `
	// [Output only] Information related to shared albums.
	// This field is only populated if the album is a shared album, the
	// developer created the album and the user has granted the
	// `photoslibrary.sharing` scope.
	ShareInfo *ShareInfo `json:"shareInfo,omitempty" `
	// Name of the album displayed to the user in their Google Photos account.
	// This string shouldn't be more than 500 characters.
	Title string `json:"title,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *Album) MarshalJSON() ([]byte, error) {
	type NoMethod Album
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Specifies a position in an album.
type AlbumPosition struct {
	// Type of position, for a media or enrichment item.
	Position string `json:"position,omitempty" `
	// The enrichment item to which the position is relative to.
	// Only used when position type is AFTER_ENRICHMENT_ITEM.
	RelativeEnrichmentItemId string `json:"relativeEnrichmentItemId,omitempty" `
	// The media item to which the position is relative to.
	// Only used when position type is AFTER_MEDIA_ITEM.
	RelativeMediaItemId string `json:"relativeMediaItemId,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *AlbumPosition) MarshalJSON() ([]byte, error) {
	type NoMethod AlbumPosition
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Request to add media items to an album.
type BatchAddMediaItemsToAlbumRequest struct {
	// Identifiers of the MediaItems to be
	// added.
	// The maximum number of media items that can be added in one call is 50.
	MediaItemIds []string `json:"mediaItemIds,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *BatchAddMediaItemsToAlbumRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchAddMediaItemsToAlbumRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Response for adding media items to an album.
type BatchAddMediaItemsToAlbumResponse struct {

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *BatchAddMediaItemsToAlbumResponse) MarshalJSON() ([]byte, error) {
	type NoMethod BatchAddMediaItemsToAlbumResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Request to create one or more media items in a user's Google Photos library.
// If an `albumid` is specified, the media items are also added to that album.
// `albumPosition` is optional and can only be specified if an `albumId` is set.
type BatchCreateMediaItemsRequest struct {
	// Identifier of the album where the media items are added. The media items
	// are also added to the user's library. This is an optional field.
	AlbumId string `json:"albumId,omitempty" `
	// Position in the album where the media items are added. If not
	// specified, the media items are added to the end of the album (as per
	// the default value, that is, `LAST_IN_ALBUM`). The request fails if this
	// field is set and the `albumId` is not specified. The request will also fail
	// if you set the field and are not the owner of the shared album.
	AlbumPosition *AlbumPosition `json:"albumPosition,omitempty" `
	// List of media items to be created.
	NewMediaItems []*NewMediaItem `json:"newMediaItems,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *BatchCreateMediaItemsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchCreateMediaItemsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// List of media items created.
type BatchCreateMediaItemsResponse struct {
	// Output only. List of media items created.
	NewMediaItemResults []*NewMediaItemResult `json:"newMediaItemResults,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *BatchCreateMediaItemsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod BatchCreateMediaItemsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Response to retrieve a list of media items.
type BatchGetMediaItemsResponse struct {
	// Output only. List of media items retrieved.
	// Note that even if the call to BatchGetMediaItems succeeds, there may have
	// been failures for some media items in the batch. These failures are
	// indicated in each
	// MediaItemResult.status.
	MediaItemResults []*MediaItemResult `json:"mediaItemResults,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *BatchGetMediaItemsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod BatchGetMediaItemsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Request to remove a list of media items from an album.
type BatchRemoveMediaItemsFromAlbumRequest struct {
	// Identifiers of the MediaItems to be
	// removed.
	//
	// Must not contain repeated identifiers and cannot be empty. The maximum
	// number of media items that can be removed in one call is 50.
	MediaItemIds []string `json:"mediaItemIds,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *BatchRemoveMediaItemsFromAlbumRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchRemoveMediaItemsFromAlbumRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Response for successfully removing all specified media items from the album.
type BatchRemoveMediaItemsFromAlbumResponse struct {

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *BatchRemoveMediaItemsFromAlbumResponse) MarshalJSON() ([]byte, error) {
	type NoMethod BatchRemoveMediaItemsFromAlbumResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// This filter allows you to return media items based on the content type.
//
// It's possible to specify a list of categories to include, and/or a list of
// categories to exclude. Within each list, the categories are combined with an
// OR. <p>
// The content filter `includedContentCategories`: [c1, c2, c3] would get media
// items that contain (c1 OR c2 OR c3). <p>
// The content filter `excludedContentCategories`: [c1, c2, c3] would NOT get
// media items that contain (c1 OR c2 OR c3). <p>
// You can also include some categories while excluding others, as in this
// example: `includedContentCategories`: [c1, c2], `excludedContentCategories`:
// [c3, c4] <p>
// The previous example would get media items that contain (c1 OR c2) AND NOT
// (c3 OR c4). A category that appears in `includedContentategories` must not
// appear in `excludedContentCategories`.
type ContentFilter struct {
	// The set of categories which are not to be included in the media item search
	// results. The items in the set are ORed. There's a maximum of 10
	// `excludedContentCategories` per request.
	ExcludedContentCategories []string `json:"excludedContentCategories,omitempty" `
	// The set of categories to be included in the media item search results.
	// The items in the set are ORed. There's a maximum of 10
	// `includedContentCategories` per request.
	IncludedContentCategories []string `json:"includedContentCategories,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *ContentFilter) MarshalJSON() ([]byte, error) {
	type NoMethod ContentFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Information about the user who added the media item. Note that this
// information is included only if the media item is within a shared album
// created by your app and you have the sharing scope.
type ContributorInfo struct {
	// Display name of the contributor.
	DisplayName string `json:"displayName,omitempty" `
	// URL to the profile picture of the contributor.
	ProfilePictureBaseUrl string `json:"profilePictureBaseUrl,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *ContributorInfo) MarshalJSON() ([]byte, error) {
	type NoMethod ContributorInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Request to create an album in Google Photos.
type CreateAlbumRequest struct {
	// The album to be created.
	Album *Album `json:"album,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *CreateAlbumRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateAlbumRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Represents a whole calendar date. The day may be 0 to represent a year and month where the day isn't significant, such as a whole calendar month. The month may be 0 to represent a a day and a year where the month isn't signficant, like when you want to specify the same day in every month of a year or a specific year. The year may be 0 to represent a month and day independent of year, like an anniversary date.
type Date struct {
	// Day of month. Must be from 1 to 31 and valid for the year and month, or 0 if specifying a year/month where the day isn't significant.
	Day int64 `json:"day,omitempty" `
	// Month of year. Must be from 1 to 12, or 0 if specifying a year without a
	// month and day.
	Month int64 `json:"month,omitempty" `
	// Year of date. Must be from 1 to 9999, or 0 if specifying a date without
	// a year.
	Year int64 `json:"year,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *Date) MarshalJSON() ([]byte, error) {
	type NoMethod Date
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// This filter defines the allowed dates or date ranges for the media returned.
// It's possible to pick a set of specific dates and a set of date ranges.
type DateFilter struct {
	// List of dates that match the media items' creation date. A maximum of
	// 5 dates can be included per request.
	Dates []*Date `json:"dates,omitempty" `
	// List of dates ranges that match the media items' creation date. A
	// maximum of 5 dates ranges can be included per request.
	Ranges []*DateRange `json:"ranges,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *DateFilter) MarshalJSON() ([]byte, error) {
	type NoMethod DateFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Defines a range of dates. Both dates must be of the same format. For more
// information, see <a href="#Date">Date</a>
type DateRange struct {
	// The end date (included as part of the range). It must be specified in the
	// same format as the start date.
	EndDate *Date `json:"endDate,omitempty" `
	// The start date (included as part of the range) in one of the formats
	// described.
	StartDate *Date `json:"startDate,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *DateRange) MarshalJSON() ([]byte, error) {
	type NoMethod DateRange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// An enrichment item.
type EnrichmentItem struct {
	// Identifier of the enrichment item.
	Id string `json:"id,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *EnrichmentItem) MarshalJSON() ([]byte, error) {
	type NoMethod EnrichmentItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// This filter defines the features that the media items should have.
type FeatureFilter struct {
	// The set of features to be included in the media item search results.
	// The items in the set are ORed and may match any of the specified features.
	IncludedFeatures []string `json:"includedFeatures,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *FeatureFilter) MarshalJSON() ([]byte, error) {
	type NoMethod FeatureFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Filters that can be applied to a media item search.
// If multiple filter options are specified, they're treated as AND with each
// other.
type Filters struct {
	// Filters the media items based on their content.
	ContentFilter *ContentFilter `json:"contentFilter,omitempty" `
	// Filters the media items based on their creation date.
	DateFilter *DateFilter `json:"dateFilter,omitempty" `
	// If set, the results exclude media items that were not created by this app.
	// Defaults to false (all media items are returned). This field is ignored if
	// the photoslibrary.readonly.appcreateddata scope is used.
	ExcludeNonAppCreatedData bool `json:"excludeNonAppCreatedData,omitempty" `
	// Filters the media items based on their features.
	FeatureFilter *FeatureFilter `json:"featureFilter,omitempty" `
	// If set, the results include media items that the user has archived.
	// Defaults to false (archived media items aren't included).
	IncludeArchivedMedia bool `json:"includeArchivedMedia,omitempty" `
	// Filters the media items based on the type of media.
	MediaTypeFilter *MediaTypeFilter `json:"mediaTypeFilter,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *Filters) MarshalJSON() ([]byte, error) {
	type NoMethod Filters
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Request to join a shared album on behalf of the user. This uses a shareToken
// which can be acquired via the shareAlbum or listSharedAlbums calls.
type JoinSharedAlbumRequest struct {
	// Token to join the shared album on behalf of the user.
	ShareToken string `json:"shareToken,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *JoinSharedAlbumRequest) MarshalJSON() ([]byte, error) {
	type NoMethod JoinSharedAlbumRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Response to successfully joining the shared album on behalf of the user.
type JoinSharedAlbumResponse struct {
	// Shared album that the user has joined.
	Album *Album `json:"album,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *JoinSharedAlbumResponse) MarshalJSON() ([]byte, error) {
	type NoMethod JoinSharedAlbumResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// An object representing a latitude/longitude pair. This is expressed as a pair
// of doubles representing degrees latitude and degrees longitude. Unless
// specified otherwise, this must conform to the
// <a href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// standard</a>. Values must be within normalized ranges.
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `json:"latitude,omitempty" `
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `json:"longitude,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *LatLng) MarshalJSON() ([]byte, error) {
	type NoMethod LatLng
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Request to leave a shared album on behalf of the user. This uses a shareToken
// which can be acquired via the or listSharedAlbums or getAlbum calls.
type LeaveSharedAlbumRequest struct {
	// Token to leave the shared album on behalf of the user.
	ShareToken string `json:"shareToken,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *LeaveSharedAlbumRequest) MarshalJSON() ([]byte, error) {
	type NoMethod LeaveSharedAlbumRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Response to successfully leaving the shared album on behalf of the user.
type LeaveSharedAlbumResponse struct {

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *LeaveSharedAlbumResponse) MarshalJSON() ([]byte, error) {
	type NoMethod LeaveSharedAlbumResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// List of albums requested.
type ListAlbumsResponse struct {
	// Output only. List of albums shown in the Albums tab of the user's Google
	// Photos app.
	Albums []*Album `json:"albums,omitempty" `
	// Output only. Token to use to get the next set of albums. Populated if
	// there are more albums to retrieve for this request.
	NextPageToken string `json:"nextPageToken,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *ListAlbumsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListAlbumsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// List of all media items from the user's Google Photos library.
type ListMediaItemsResponse struct {
	// Output only. List of media items in the user's library.
	MediaItems []*MediaItem `json:"mediaItems,omitempty" `
	// Output only. Token to use to get the next set of media items. Its presence
	// is the only reliable indicator of more media items being available in the
	// next request.
	NextPageToken string `json:"nextPageToken,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *ListMediaItemsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListMediaItemsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// List of shared albums requested.
type ListSharedAlbumsResponse struct {
	// Output only. Token to use to get the next set of shared albums. Populated
	// if there are more shared albums to retrieve for this request.
	NextPageToken string `json:"nextPageToken,omitempty" `
	// Output only. List of shared albums.
	SharedAlbums []*Album `json:"sharedAlbums,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *ListSharedAlbumsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListSharedAlbumsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Represents a physical location.
type Location struct {
	// Position of the location on the map.
	Latlng *LatLng `json:"latlng,omitempty" `
	// Name of the location to be displayed.
	LocationName string `json:"locationName,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *Location) MarshalJSON() ([]byte, error) {
	type NoMethod Location
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// An enrichment containing a single location.
type LocationEnrichment struct {
	// Location for this enrichment item.
	Location *Location `json:"location,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *LocationEnrichment) MarshalJSON() ([]byte, error) {
	type NoMethod LocationEnrichment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// An enrichment containing a map, showing origin and destination locations.
type MapEnrichment struct {
	// Destination location for this enrichemt item.
	Destination *Location `json:"destination,omitempty" `
	// Origin location for this enrichment item.
	Origin *Location `json:"origin,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *MapEnrichment) MarshalJSON() ([]byte, error) {
	type NoMethod MapEnrichment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Representation of a media item (such as a photo or video) in Google Photos.
type MediaItem struct {
	// A URL to the media item's bytes. This shouldn't be used as is. Parameters
	// should be appended to this URL before use. See the [developer
	// documentation](https://developers.google.com/photos/library/guides/access-media-items#base-urls)
	// for a complete list of supported parameters. For example, `'=w2048-h1024'`
	// will set the dimensions of a media item of type photo to have a width of
	// 2048 px and height of 1024 px.
	BaseUrl string `json:"baseUrl,omitempty" `
	// Information about the user who created this media item.
	ContributorInfo *ContributorInfo `json:"contributorInfo,omitempty" `
	// Description of the media item. This is shown to the user in the item's
	// info section in the Google Photos app.
	Description string `json:"description,omitempty" `
	// Filename of the media item. This is shown to the user in the item's info
	// section in the Google Photos app.
	Filename string `json:"filename,omitempty" `
	// Identifier for the media item. This is a persistent identifier that can be
	// used between sessions to identify this media item.
	Id string `json:"id,omitempty" `
	// Metadata related to the media item, such as, height, width, or
	// creation time.
	MediaMetadata *MediaMetadata `json:"mediaMetadata,omitempty" `
	// MIME type of the media item. For example, `image/jpeg`.
	MimeType string `json:"mimeType,omitempty" `
	// Google Photos URL for the media item. This link is available to
	// the user only if they're signed in.
	ProductUrl string `json:"productUrl,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *MediaItem) MarshalJSON() ([]byte, error) {
	type NoMethod MediaItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Result of retrieving a media item.
type MediaItemResult struct {
	// Media item retrieved from the user's library. It's populated if no errors
	// occurred and the media item was fetched successfully.
	MediaItem *MediaItem `json:"mediaItem,omitempty" `
	// If an error occurred while accessing this media item, this field
	// is populated with information related to the error. For details regarding
	// this field, see Status.
	Status *Status `json:"status,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *MediaItemResult) MarshalJSON() ([]byte, error) {
	type NoMethod MediaItemResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Metadata for a media item.
type MediaMetadata struct {
	// Time when the media item was first created (not when it was uploaded to
	// Google Photos).
	CreationTime string `json:"creationTime,omitempty" `
	// Original height (in pixels) of the media item.
	Height string `json:"height,omitempty" `
	// Metadata for a photo media type.
	Photo *Photo `json:"photo,omitempty" `
	// Metadata for a video media type.
	Video *Video `json:"video,omitempty" `
	// Original width (in pixels) of the media item.
	Width string `json:"width,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *MediaMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod MediaMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// This filter defines the type of media items to be returned, for example,
// videos or photos. All the specified media types are treated as an OR when
// used together.
type MediaTypeFilter struct {
	// The types of media items to be included. This field should be populated
	// with only one media type. If you specify multiple media types, it results
	// in an error.
	MediaTypes []string `json:"mediaTypes,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *MediaTypeFilter) MarshalJSON() ([]byte, error) {
	type NoMethod MediaTypeFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// A new enrichment item to be added to an album, used by the
// `albums.addEnrichment` call.
type NewEnrichmentItem struct {
	// Location to be added to the album.
	LocationEnrichment *LocationEnrichment `json:"locationEnrichment,omitempty" `
	// Map to be added to the album.
	MapEnrichment *MapEnrichment `json:"mapEnrichment,omitempty" `
	// Text to be added to the album.
	TextEnrichment *TextEnrichment `json:"textEnrichment,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *NewEnrichmentItem) MarshalJSON() ([]byte, error) {
	type NoMethod NewEnrichmentItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// New media item that's created in a user's Google Photos account.
type NewMediaItem struct {
	// Description of the media item. This will be shown to the user in the item's
	// info section in the Google Photos app.
	// This string shouldn't be more than 1000 characters.
	Description string `json:"description,omitempty" `
	// A new media item that has been uploaded via the included `uploadToken`.
	SimpleMediaItem *SimpleMediaItem `json:"simpleMediaItem,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *NewMediaItem) MarshalJSON() ([]byte, error) {
	type NoMethod NewMediaItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Result of creating a new media item.
type NewMediaItemResult struct {
	// Media item created with the upload token. It's populated if no errors
	// occurred and the media item was created successfully.
	MediaItem *MediaItem `json:"mediaItem,omitempty" `
	// If an error occurred during the creation of this media item, this field
	// is  populated with information related to the error. For details regarding
	// this field, see <a href="#Status">Status</a>.
	Status *Status `json:"status,omitempty" `
	// The upload token used to create this new media item.
	UploadToken string `json:"uploadToken,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *NewMediaItemResult) MarshalJSON() ([]byte, error) {
	type NoMethod NewMediaItemResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Metadata that is specific to a photo, such as, ISO, focal length and
// exposure time. Some of these fields may be null or not included.
type Photo struct {
	// Aperture f number of the camera lens with which the photo was taken.
	ApertureFNumber float64 `json:"apertureFNumber,omitempty" `
	// Brand of the camera with which the photo was taken.
	CameraMake string `json:"cameraMake,omitempty" `
	// Model of the camera with which the photo was taken.
	CameraModel string `json:"cameraModel,omitempty" `
	// Exposure time of the camera aperture when the photo was taken.
	ExposureTime string `json:"exposureTime,omitempty" `
	// Focal length of the camera lens with which the photo was taken.
	FocalLength float64 `json:"focalLength,omitempty" `
	// ISO of the camera with which the photo was taken.
	IsoEquivalent int64 `json:"isoEquivalent,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *Photo) MarshalJSON() ([]byte, error) {
	type NoMethod Photo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Request to search for media items in a user's library.
//
// If the album id is specified, this call will return the list of media items
// in the album. If neither filters nor album id are
// specified, this call will return all media items in a user's Google Photos
// library.
//
// If filters are specified, this call will return all media items in
// the user's library that fulfill the filter criteria.
//
// Filters and album id must not both be set, as this will result in an
// invalid request.
type SearchMediaItemsRequest struct {
	// Identifier of an album. If populated, lists all media items in
	// specified album. Can't set in conjunction with any filters.
	AlbumId string `json:"albumId,omitempty" `
	// Filters to apply to the request. Can't be set in conjunction with an
	// `albumId`.
	Filters *Filters `json:"filters,omitempty" `
	// Maximum number of media items to return in the response. The default number
	// of media items to return at a time is 25. The maximum
	// `pageSize` is 100.
	PageSize int64 `json:"pageSize,omitempty" `
	// A continuation token to get the next page of the results. Adding this to
	// the request returns the rows after the `pageToken`. The `pageToken` should
	// be the value returned in the `nextPageToken` parameter in the response to
	// the `searchMediaItems` request.
	PageToken string `json:"pageToken,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *SearchMediaItemsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SearchMediaItemsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// List of media items that match the search parameters.
type SearchMediaItemsResponse struct {
	// Output only. List of media items that match the search parameters.
	MediaItems []*MediaItem `json:"mediaItems,omitempty" `
	// Output only. Use this token to get the next set of media items. Its
	// presence is the only reliable indicator of more media items being available
	// in the next request.
	NextPageToken string `json:"nextPageToken,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *SearchMediaItemsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod SearchMediaItemsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Request to make an album shared in Google Photos.
type ShareAlbumRequest struct {
	// Options to be set when converting the album to a shared album.
	SharedAlbumOptions *SharedAlbumOptions `json:"sharedAlbumOptions,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *ShareAlbumRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ShareAlbumRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Response to successfully sharing an album.
type ShareAlbumResponse struct {
	// Output only. Information about the shared album.
	ShareInfo *ShareInfo `json:"shareInfo,omitempty" `

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *ShareAlbumResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ShareAlbumResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Information about albums that are shared. This information is only included
// if you created the album, it is shared and you have the sharing scope.
type ShareInfo struct {
	// True if the user has joined the album. This is always true for the owner
	// of the shared album.
	IsJoined bool `json:"isJoined,omitempty" `
	// True if the user owns the album.
	IsOwned bool `json:"isOwned,omitempty" `
	// A token that can be used by other users to join this shared album via the
	// API.
	ShareToken string `json:"shareToken,omitempty" `
	// A link to the album that's now shared on the Google Photos website and app.
	// Anyone with the link can access this shared album and see all of the items
	// present in the album.
	ShareableUrl string `json:"shareableUrl,omitempty" `
	// Options that control the sharing of an album.
	SharedAlbumOptions *SharedAlbumOptions `json:"sharedAlbumOptions,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *ShareInfo) MarshalJSON() ([]byte, error) {
	type NoMethod ShareInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Options that control the sharing of an album.
type SharedAlbumOptions struct {
	// True if the shared album allows collaborators (users who have joined
	// the album) to add media items to it. Defaults to false.
	IsCollaborative bool `json:"isCollaborative,omitempty" `
	// True if the shared album allows the owner and the collaborators (users
	// who have joined the album) to add comments to the album. Defaults to false.
	IsCommentable bool `json:"isCommentable,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *SharedAlbumOptions) MarshalJSON() ([]byte, error) {
	type NoMethod SharedAlbumOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// A simple media item to be created in Google Photos via an upload token.
type SimpleMediaItem struct {
	// Token identifying the media bytes that have been uploaded to Google.
	UploadToken string `json:"uploadToken,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *SimpleMediaItem) MarshalJSON() ([]byte, error) {
	type NoMethod SimpleMediaItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// The `Status` type defines a logical error model that is suitable for
// different programming environments, including REST APIs and RPC APIs. It is
// used by [gRPC](https://github.com/grpc). Each `Status` message contains
// three pieces of data: error code, error message, and error details.
//
// You can find out more about this error model and how to work with it in the
// [API Design Guide](https://cloud.google.com/apis/design/errors).
type Status struct {
	// The status code, which should be an enum value of google.rpc.Code.
	Code int64 `json:"code,omitempty" `
	// A list of messages that carry the error details.  There is a common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty" `
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// An enrichment containing text.
type TextEnrichment struct {
	// Text for this enrichment item.
	Text string `json:"text,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *TextEnrichment) MarshalJSON() ([]byte, error) {
	type NoMethod TextEnrichment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Request to unshare a shared album in Google Photos.
type UnshareAlbumRequest struct {
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *UnshareAlbumRequest) MarshalJSON() ([]byte, error) {
	type NoMethod UnshareAlbumRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Response of a successful unshare of a shared album.
type UnshareAlbumResponse struct {

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

func (s *UnshareAlbumResponse) MarshalJSON() ([]byte, error) {
	type NoMethod UnshareAlbumResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Metadata that is specific to a video, for example, fps and processing status.
// Some of these fields may be null or not included.
type Video struct {
	// Brand of the camera with which the video was taken.
	CameraMake string `json:"cameraMake,omitempty" `
	// Model of the camera with which the video was taken.
	CameraModel string `json:"cameraModel,omitempty" `
	// Frame rate of the video.
	Fps float64 `json:"fps,omitempty" `
	// Processing status of the video.
	Status string `json:"status,omitempty" `

	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

func (s *Video) MarshalJSON() ([]byte, error) {
	type NoMethod Video
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "photoslibrary.albums.addEnrichment":

type AlbumsAddEnrichmentCall struct {
	s                           *Service
	albumId                     string
	addenrichmenttoalbumrequest *AddEnrichmentToAlbumRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// addEnrichment: // Adds an enrichment at a specified position in a defined album.
func (r *AlbumsService) AddEnrichment(albumId string, addenrichmenttoalbumrequest *AddEnrichmentToAlbumRequest) *AlbumsAddEnrichmentCall {
	c := &AlbumsAddEnrichmentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.albumId = albumId
	c.addenrichmenttoalbumrequest = addenrichmenttoalbumrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlbumsAddEnrichmentCall) Fields(s ...googleapi.Field) *AlbumsAddEnrichmentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlbumsAddEnrichmentCall) Context(ctx context.Context) *AlbumsAddEnrichmentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlbumsAddEnrichmentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlbumsAddEnrichmentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.addenrichmenttoalbumrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/albums/{+albumId}:addEnrichment")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"albumId": fmt.Sprintf("%v", c.albumId),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.albums.addEnrichment" call.
// Exactly one of *AddEnrichmentToAlbumResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *AddEnrichmentToAlbumResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlbumsAddEnrichmentCall) Do(opts ...googleapi.CallOption) (*AddEnrichmentToAlbumResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AddEnrichmentToAlbumResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.albums.batchAddMediaItems":

type AlbumsBatchAddMediaItemsCall struct {
	s                                *Service
	albumId                          string
	batchaddmediaitemstoalbumrequest *BatchAddMediaItemsToAlbumRequest
	urlParams_                       gensupport.URLParams
	ctx_                             context.Context
	header_                          http.Header
}

// batchAddMediaItems: // Adds one or more media items in a user's Google Photos library to
// an album. The media items and albums must have been created by the
// developer via the API.
//
// Media items are added to the end of the album. If multiple media items are
// given, they are added in the order specified in this call.
//
// Each album can contain up to 20,000 media items.
//
// Only media items that are in the user's library can be added to an
// album. For albums that are shared, the album must either be owned by the
// user or the user must have joined the album as a collaborator.
//
// Partial success is not supported. The entire request will fail if an
// invalid media item or album is specified.
func (r *AlbumsService) BatchAddMediaItems(albumId string, batchaddmediaitemstoalbumrequest *BatchAddMediaItemsToAlbumRequest) *AlbumsBatchAddMediaItemsCall {
	c := &AlbumsBatchAddMediaItemsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.albumId = albumId
	c.batchaddmediaitemstoalbumrequest = batchaddmediaitemstoalbumrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlbumsBatchAddMediaItemsCall) Fields(s ...googleapi.Field) *AlbumsBatchAddMediaItemsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlbumsBatchAddMediaItemsCall) Context(ctx context.Context) *AlbumsBatchAddMediaItemsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlbumsBatchAddMediaItemsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlbumsBatchAddMediaItemsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchaddmediaitemstoalbumrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/albums/{+albumId}:batchAddMediaItems")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"albumId": fmt.Sprintf("%v", c.albumId),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.albums.batchAddMediaItems" call.
// Exactly one of *BatchAddMediaItemsToAlbumResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *BatchAddMediaItemsToAlbumResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlbumsBatchAddMediaItemsCall) Do(opts ...googleapi.CallOption) (*BatchAddMediaItemsToAlbumResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BatchAddMediaItemsToAlbumResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.albums.batchRemoveMediaItems":

type AlbumsBatchRemoveMediaItemsCall struct {
	s                                     *Service
	albumId                               string
	batchremovemediaitemsfromalbumrequest *BatchRemoveMediaItemsFromAlbumRequest
	urlParams_                            gensupport.URLParams
	ctx_                                  context.Context
	header_                               http.Header
}

// batchRemoveMediaItems: // Removes one or more media items from a specified album. The media items and
// the album must have been created by the developer via the API.
//
// For albums that are shared, this action is only supported for media items
// that were added to the album by this user, or for all media items if the
// album was created by this user.
//
// Partial success is not supported. The entire request will fail and no
// action will be performed on the album if an invalid media item or album is
// specified.
func (r *AlbumsService) BatchRemoveMediaItems(albumId string, batchremovemediaitemsfromalbumrequest *BatchRemoveMediaItemsFromAlbumRequest) *AlbumsBatchRemoveMediaItemsCall {
	c := &AlbumsBatchRemoveMediaItemsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.albumId = albumId
	c.batchremovemediaitemsfromalbumrequest = batchremovemediaitemsfromalbumrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlbumsBatchRemoveMediaItemsCall) Fields(s ...googleapi.Field) *AlbumsBatchRemoveMediaItemsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlbumsBatchRemoveMediaItemsCall) Context(ctx context.Context) *AlbumsBatchRemoveMediaItemsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlbumsBatchRemoveMediaItemsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlbumsBatchRemoveMediaItemsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchremovemediaitemsfromalbumrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/albums/{+albumId}:batchRemoveMediaItems")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"albumId": fmt.Sprintf("%v", c.albumId),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.albums.batchRemoveMediaItems" call.
// Exactly one of *BatchRemoveMediaItemsFromAlbumResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *BatchRemoveMediaItemsFromAlbumResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlbumsBatchRemoveMediaItemsCall) Do(opts ...googleapi.CallOption) (*BatchRemoveMediaItemsFromAlbumResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BatchRemoveMediaItemsFromAlbumResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.albums.create":

type AlbumsCreateCall struct {
	s                  *Service
	createalbumrequest *CreateAlbumRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// create: // Creates an album in a user's Google Photos library.
func (r *AlbumsService) Create(createalbumrequest *CreateAlbumRequest) *AlbumsCreateCall {
	c := &AlbumsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.createalbumrequest = createalbumrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlbumsCreateCall) Fields(s ...googleapi.Field) *AlbumsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlbumsCreateCall) Context(ctx context.Context) *AlbumsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlbumsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlbumsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createalbumrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/albums")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders

	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.albums.create" call.
// Exactly one of *Album or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *Album.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlbumsCreateCall) Do(opts ...googleapi.CallOption) (*Album, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Album{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.albums.get":

type AlbumsGetCall struct {
	s            *Service
	albumId      string
	ifNoneMatch_ string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// get: // Returns the album based on the specified `albumId`.
// The `albumId` must be the ID of an album owned by the user or a shared
// album that the user has joined.
func (r *AlbumsService) Get(albumId string) *AlbumsGetCall {
	c := &AlbumsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.albumId = albumId

	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlbumsGetCall) Fields(s ...googleapi.Field) *AlbumsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlbumsGetCall) Context(ctx context.Context) *AlbumsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlbumsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlbumsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/albums/{+albumId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"albumId": fmt.Sprintf("%v", c.albumId),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.albums.get" call.
// Exactly one of *Album or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *Album.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlbumsGetCall) Do(opts ...googleapi.CallOption) (*Album, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Album{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.albums.list":

type AlbumsListCall struct {
	s                        *Service
	excludeNonAppCreatedData bool
	ifNoneMatch_             string
	urlParams_               gensupport.URLParams
	ctx_                     context.Context
	header_                  http.Header
}

// list: // Lists all albums shown to a user in the Albums tab of the Google
// Photos app.
func (r *AlbumsService) List(excludeNonAppCreatedData bool) *AlbumsListCall {
	c := &AlbumsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.excludeNonAppCreatedData = excludeNonAppCreatedData

	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlbumsListCall) Fields(s ...googleapi.Field) *AlbumsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlbumsListCall) Context(ctx context.Context) *AlbumsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlbumsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// albums to return in the response. The default number of
// albums to return at a time is 20. The maximum page size is 50.
func (c *AlbumsListCall) PageSize(pageSize int64) *AlbumsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token to get the next page of the results. Adding this to
// the request will return the rows after the pageToken. The pageToken
// should
// be the value returned in the nextPageToken parameter in the response
// to the
// listSharedAlbums request.
func (c *AlbumsListCall) PageToken(pageToken string) *AlbumsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AlbumsListCall) Pages(ctx context.Context, f func(*ListAlbumsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

func (c *AlbumsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/albums")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"excludeNonAppCreatedData": fmt.Sprintf("%v", c.excludeNonAppCreatedData),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.albums.list" call.
// Exactly one of *ListAlbumsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListAlbumsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlbumsListCall) Do(opts ...googleapi.CallOption) (*ListAlbumsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListAlbumsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.albums.share":

type AlbumsShareCall struct {
	s                 *Service
	albumId           string
	sharealbumrequest *ShareAlbumRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// share: // Marks an album as shared and accessible to other users. This action can
// only be performed on albums which were created by the developer via the
// API.
func (r *AlbumsService) Share(albumId string, sharealbumrequest *ShareAlbumRequest) *AlbumsShareCall {
	c := &AlbumsShareCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.albumId = albumId
	c.sharealbumrequest = sharealbumrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlbumsShareCall) Fields(s ...googleapi.Field) *AlbumsShareCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlbumsShareCall) Context(ctx context.Context) *AlbumsShareCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlbumsShareCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlbumsShareCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.sharealbumrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/albums/{+albumId}:share")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"albumId": fmt.Sprintf("%v", c.albumId),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.albums.share" call.
// Exactly one of *ShareAlbumResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ShareAlbumResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlbumsShareCall) Do(opts ...googleapi.CallOption) (*ShareAlbumResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ShareAlbumResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.albums.unshare":

type AlbumsUnshareCall struct {
	s                   *Service
	albumId             string
	unsharealbumrequest *UnshareAlbumRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// unshare: // Marks a previously shared album as private. This means that the album is
// no longer shared and all the non-owners will lose access to the album. All
// non-owner content will be removed from the album. If a non-owner has
// previously added the album to their library, they will retain all photos in
// their library. This action can only be performed on albums which were
// created by the developer via the API.
func (r *AlbumsService) Unshare(albumId string, unsharealbumrequest *UnshareAlbumRequest) *AlbumsUnshareCall {
	c := &AlbumsUnshareCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.albumId = albumId
	c.unsharealbumrequest = unsharealbumrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AlbumsUnshareCall) Fields(s ...googleapi.Field) *AlbumsUnshareCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AlbumsUnshareCall) Context(ctx context.Context) *AlbumsUnshareCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AlbumsUnshareCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AlbumsUnshareCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.unsharealbumrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/albums/{+albumId}:unshare")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"albumId": fmt.Sprintf("%v", c.albumId),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.albums.unshare" call.
// Exactly one of *UnshareAlbumResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *UnshareAlbumResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AlbumsUnshareCall) Do(opts ...googleapi.CallOption) (*UnshareAlbumResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UnshareAlbumResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.mediaItems.batchCreate":

type MediaItemsBatchCreateCall struct {
	s                            *Service
	batchcreatemediaitemsrequest *BatchCreateMediaItemsRequest
	urlParams_                   gensupport.URLParams
	ctx_                         context.Context
	header_                      http.Header
}

// batchCreate: // Creates one or more media items in a user's Google Photos library.
//
// This is the second step for creating a media item. For details regarding
// Step 1, uploading the raw bytes to a Google Server, see
// <a href="/photos/library/guides/upload-media">Uploading media</a>.
//
// This call adds the media item to the library. If an album `id` is
// specified, the call adds the media item to the album too. Each album can
// contain up to 20,000 media items. By default, the media item will be added
// to the end of the library or album.
//
// If an album `id` and position are both defined, the media item is
// added to the album at the specified position.
//
// If the call contains multiple media items, they're added at the specified
// position.
// If you are creating a media item in a shared album where you are not the
// owner, you are not allowed to position the media item. Doing so will result
// in a `BAD REQUEST` error.
func (r *MediaItemsService) BatchCreate(batchcreatemediaitemsrequest *BatchCreateMediaItemsRequest) *MediaItemsBatchCreateCall {
	c := &MediaItemsBatchCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.batchcreatemediaitemsrequest = batchcreatemediaitemsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MediaItemsBatchCreateCall) Fields(s ...googleapi.Field) *MediaItemsBatchCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MediaItemsBatchCreateCall) Context(ctx context.Context) *MediaItemsBatchCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MediaItemsBatchCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MediaItemsBatchCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchcreatemediaitemsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/mediaItems:batchCreate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders

	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.mediaItems.batchCreate" call.
// Exactly one of *BatchCreateMediaItemsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *BatchCreateMediaItemsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MediaItemsBatchCreateCall) Do(opts ...googleapi.CallOption) (*BatchCreateMediaItemsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BatchCreateMediaItemsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.mediaItems.batchGet":

type MediaItemsBatchGetCall struct {
	s            *Service
	mediaItemIds string
	ifNoneMatch_ string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// batchGet: // Returns the list of media items for the specified media item identifiers.
// Items are returned in the same order as the supplied identifiers.
func (r *MediaItemsService) BatchGet(mediaItemIds string) *MediaItemsBatchGetCall {
	c := &MediaItemsBatchGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.mediaItemIds = mediaItemIds

	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MediaItemsBatchGetCall) Fields(s ...googleapi.Field) *MediaItemsBatchGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MediaItemsBatchGetCall) Context(ctx context.Context) *MediaItemsBatchGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MediaItemsBatchGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MediaItemsBatchGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/mediaItems:batchGet")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"mediaItemIds": fmt.Sprintf("%v", c.mediaItemIds),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.mediaItems.batchGet" call.
// Exactly one of *BatchGetMediaItemsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *BatchGetMediaItemsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MediaItemsBatchGetCall) Do(opts ...googleapi.CallOption) (*BatchGetMediaItemsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BatchGetMediaItemsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.mediaItems.get":

type MediaItemsGetCall struct {
	s            *Service
	mediaItemId  string
	ifNoneMatch_ string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// get: // Returns the media item for the specified media item identifier.
func (r *MediaItemsService) Get(mediaItemId string) *MediaItemsGetCall {
	c := &MediaItemsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.mediaItemId = mediaItemId

	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MediaItemsGetCall) Fields(s ...googleapi.Field) *MediaItemsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MediaItemsGetCall) Context(ctx context.Context) *MediaItemsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MediaItemsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MediaItemsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/mediaItems/{+mediaItemId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"mediaItemId": fmt.Sprintf("%v", c.mediaItemId),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.mediaItems.get" call.
// Exactly one of *MediaItem or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *MediaItem.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MediaItemsGetCall) Do(opts ...googleapi.CallOption) (*MediaItem, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &MediaItem{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.mediaItems.list":

type MediaItemsListCall struct {
	s            *Service
	ifNoneMatch_ string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// list: // List all media items from a user's Google Photos library.
func (r *MediaItemsService) List() *MediaItemsListCall {
	c := &MediaItemsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}

	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MediaItemsListCall) Fields(s ...googleapi.Field) *MediaItemsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MediaItemsListCall) Context(ctx context.Context) *MediaItemsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MediaItemsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// albums to return in the response. The default number of
// albums to return at a time is 20. The maximum page size is 50.
func (c *MediaItemsListCall) PageSize(pageSize int64) *MediaItemsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token to get the next page of the results. Adding this to
// the request will return the rows after the pageToken. The pageToken
// should
// be the value returned in the nextPageToken parameter in the response
// to the
// listSharedAlbums request.
func (c *MediaItemsListCall) PageToken(pageToken string) *MediaItemsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *MediaItemsListCall) Pages(ctx context.Context, f func(*ListMediaItemsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

func (c *MediaItemsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/mediaItems")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders

	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.mediaItems.list" call.
// Exactly one of *ListMediaItemsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListMediaItemsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MediaItemsListCall) Do(opts ...googleapi.CallOption) (*ListMediaItemsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListMediaItemsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.mediaItems.search":

type MediaItemsSearchCall struct {
	s                       *Service
	searchmediaitemsrequest *SearchMediaItemsRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
	header_                 http.Header
}

// search: // Searches for media items in a user's Google Photos library.
// If no filters are set, then all media items in the user's library are
// returned.
// If an album is set, all media items in the specified album are returned.
// If filters are specified, media items that match the filters from the
// user's library are listed. If you set both the album and the filters, the
// request results in an error.
func (r *MediaItemsService) Search(searchmediaitemsrequest *SearchMediaItemsRequest) *MediaItemsSearchCall {
	c := &MediaItemsSearchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.searchmediaitemsrequest = searchmediaitemsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MediaItemsSearchCall) Fields(s ...googleapi.Field) *MediaItemsSearchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MediaItemsSearchCall) Context(ctx context.Context) *MediaItemsSearchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MediaItemsSearchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// albums to return in the response. The default number of
// albums to return at a time is 20. The maximum page size is 50.
func (c *MediaItemsSearchCall) PageSize(pageSize int64) *MediaItemsSearchCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token to get the next page of the results. Adding this to
// the request will return the rows after the pageToken. The pageToken
// should
// be the value returned in the nextPageToken parameter in the response
// to the
// listSharedAlbums request.
func (c *MediaItemsSearchCall) PageToken(pageToken string) *MediaItemsSearchCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *MediaItemsSearchCall) Pages(ctx context.Context, f func(*SearchMediaItemsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

func (c *MediaItemsSearchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchmediaitemsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/mediaItems:search")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders

	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.mediaItems.search" call.
// Exactly one of *SearchMediaItemsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *SearchMediaItemsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MediaItemsSearchCall) Do(opts ...googleapi.CallOption) (*SearchMediaItemsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SearchMediaItemsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.sharedAlbums.get":

type SharedAlbumsGetCall struct {
	s            *Service
	shareToken   string
	ifNoneMatch_ string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// get: // Returns the album based on the specified `shareToken`.
func (r *SharedAlbumsService) Get(shareToken string) *SharedAlbumsGetCall {
	c := &SharedAlbumsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.shareToken = shareToken

	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SharedAlbumsGetCall) Fields(s ...googleapi.Field) *SharedAlbumsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SharedAlbumsGetCall) Context(ctx context.Context) *SharedAlbumsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SharedAlbumsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SharedAlbumsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/sharedAlbums/{+shareToken}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"shareToken": fmt.Sprintf("%v", c.shareToken),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.sharedAlbums.get" call.
// Exactly one of *Album or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *Album.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SharedAlbumsGetCall) Do(opts ...googleapi.CallOption) (*Album, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Album{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.sharedAlbums.join":

type SharedAlbumsJoinCall struct {
	s                      *Service
	joinsharedalbumrequest *JoinSharedAlbumRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// join: // Joins a shared album on behalf of the Google Photos user.
func (r *SharedAlbumsService) Join(joinsharedalbumrequest *JoinSharedAlbumRequest) *SharedAlbumsJoinCall {
	c := &SharedAlbumsJoinCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.joinsharedalbumrequest = joinsharedalbumrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SharedAlbumsJoinCall) Fields(s ...googleapi.Field) *SharedAlbumsJoinCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SharedAlbumsJoinCall) Context(ctx context.Context) *SharedAlbumsJoinCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SharedAlbumsJoinCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SharedAlbumsJoinCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.joinsharedalbumrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/sharedAlbums:join")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders

	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.sharedAlbums.join" call.
// Exactly one of *JoinSharedAlbumResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *JoinSharedAlbumResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SharedAlbumsJoinCall) Do(opts ...googleapi.CallOption) (*JoinSharedAlbumResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &JoinSharedAlbumResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.sharedAlbums.leave":

type SharedAlbumsLeaveCall struct {
	s                       *Service
	leavesharedalbumrequest *LeaveSharedAlbumRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
	header_                 http.Header
}

// leave: // Leaves a previously-joined shared album on behalf of the Google Photos
// user. The user must not own this album.
func (r *SharedAlbumsService) Leave(leavesharedalbumrequest *LeaveSharedAlbumRequest) *SharedAlbumsLeaveCall {
	c := &SharedAlbumsLeaveCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.leavesharedalbumrequest = leavesharedalbumrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SharedAlbumsLeaveCall) Fields(s ...googleapi.Field) *SharedAlbumsLeaveCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SharedAlbumsLeaveCall) Context(ctx context.Context) *SharedAlbumsLeaveCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SharedAlbumsLeaveCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SharedAlbumsLeaveCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.leavesharedalbumrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/sharedAlbums:leave")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders

	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.sharedAlbums.leave" call.
// Exactly one of *LeaveSharedAlbumResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *LeaveSharedAlbumResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SharedAlbumsLeaveCall) Do(opts ...googleapi.CallOption) (*LeaveSharedAlbumResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &LeaveSharedAlbumResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// method id "photoslibrary.sharedAlbums.list":

type SharedAlbumsListCall struct {
	s                        *Service
	excludeNonAppCreatedData bool
	ifNoneMatch_             string
	urlParams_               gensupport.URLParams
	ctx_                     context.Context
	header_                  http.Header
}

// list: // Lists all shared albums available in the Sharing tab of the
// user's Google Photos app.
func (r *SharedAlbumsService) List(excludeNonAppCreatedData bool) *SharedAlbumsListCall {
	c := &SharedAlbumsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.excludeNonAppCreatedData = excludeNonAppCreatedData

	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SharedAlbumsListCall) Fields(s ...googleapi.Field) *SharedAlbumsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SharedAlbumsListCall) Context(ctx context.Context) *SharedAlbumsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SharedAlbumsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// albums to return in the response. The default number of
// albums to return at a time is 20. The maximum page size is 50.
func (c *SharedAlbumsListCall) PageSize(pageSize int64) *SharedAlbumsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token to get the next page of the results. Adding this to
// the request will return the rows after the pageToken. The pageToken
// should
// be the value returned in the nextPageToken parameter in the response
// to the
// listSharedAlbums request.
func (c *SharedAlbumsListCall) PageToken(pageToken string) *SharedAlbumsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *SharedAlbumsListCall) Pages(ctx context.Context, f func(*ListSharedAlbumsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

func (c *SharedAlbumsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/sharedAlbums")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"excludeNonAppCreatedData": fmt.Sprintf("%v", c.excludeNonAppCreatedData),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "photoslibrary.sharedAlbums.list" call.
// Exactly one of *ListSharedAlbumsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListSharedAlbumsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SharedAlbumsListCall) Do(opts ...googleapi.CallOption) (*ListSharedAlbumsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListSharedAlbumsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}
