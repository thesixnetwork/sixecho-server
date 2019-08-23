"""create hash images

Revision ID: 5d519b033846
Revises: 697ad27726a8
Create Date: 2019-08-23 15:29:55.639581

"""
import sqlalchemy as sa
from sqlalchemy.sql import func

from alembic import op

# revision identifiers, used by Alembic.
revision = '5d519b033846'
down_revision = '697ad27726a8'
branch_labels = None
depends_on = None


def upgrade():
    op.create_table(
        'hash_images', sa.Column('id', sa.String(25), primary_key=True),
        sa.Column('digital_content_id', sa.Integer, nullable=False),
        sa.Column('hash_type', sa.String(10), nullable=False),
        sa.Column('created_at',
                  sa.DateTime(timezone=True),
                  nullable=False,
                  server_default=func.now()),
        sa.Column('updated_at',
                  sa.DateTime(timezone=True),
                  nullable=False,
                  server_default=func.now()),
        sa.Index("digital_content_id_index", "digital_content_id"))


def downgrade():
    op.drop_table('hash_images')
