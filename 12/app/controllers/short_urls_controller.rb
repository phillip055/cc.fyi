class ShortUrlsController < ApplicationController
  before_action :set_short_url, only: %i[ show edit update destroy ]

  # GET /short_urls or /short_urls.json
  def index
    @short_urls = ShortUrl.all
  end

  # GET /short_urls/1 or /short_urls/1.json
  def show
  end

  def redirect
    @short_url = ShortUrl.find_by(hash_key: request.path[1..-1])
    if @short_url
      redirect_to @short_url.long_url, allow_other_host: true
    else
      redirect_to "/404"
    end
  end

  # GET /short_urls/new
  def new
    @short_url = ShortUrl.new
  end

  # GET /short_urls/1/edit
  def edit
  end

  # POST /short_urls or /short_urls.json
  def create
    @short_url = ShortUrl.new(short_url_params)

    respond_to do |format|
      if @short_url.save
        format.html { redirect_to short_url_url(@short_url), notice: "Short url was successfully created." }
        format.json { render :show, status: :created, location: @short_url }
      else
        format.html { render :new, status: :unprocessable_entity }
        format.json { render json: @short_url.errors, status: :unprocessable_entity }
      end
    end
  end

  # PATCH/PUT /short_urls/1 or /short_urls/1.json
  def update
    respond_to do |format|
      if @short_url.update(short_url_params)
        format.html { redirect_to short_url_url(@short_url), notice: "Short url was successfully updated." }
        format.json { render :show, status: :ok, location: @short_url }
      else
        format.html { render :edit, status: :unprocessable_entity }
        format.json { render json: @short_url.errors, status: :unprocessable_entity }
      end
    end
  end

  # DELETE /short_urls/1 or /short_urls/1.json
  def destroy
    @short_url.destroy!

    respond_to do |format|
      format.html { redirect_to short_urls_url, notice: "Short url was successfully destroyed." }
      format.json { head :no_content }
    end
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_short_url
      @short_url = ShortUrl.find(params[:id])
    end

    # Only allow a list of trusted parameters through.
    def short_url_params
      params.require(:short_url).permit(:long_url, :hash_key)
    end
end
